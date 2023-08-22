package controllers

import (
	"github.com/keenangit/pasaman/forms"
	"github.com/keenangit/pasaman/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

//ArticleController ...
type ArticleController struct{}

var articleModel = new(models.ArticleModel)
var articleForm = new(forms.ArticleForm)

//Create ...
func (ctrl ArticleController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreateArticleForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := articleForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := articleModel.Create(userID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article created", "id": id})
}

//All ...
func (ctrl ArticleController) All(c *gin.Context) {
	// userID := getUserID(c)
	userID := "public"

	results, err := articleModel.GetAll(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get articles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

//One ...
func (ctrl ArticleController) One(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	data, err := articleModel.One(userID, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

//Update ...
func (ctrl ArticleController) Update(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	var form forms.CreateArticleForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := articleForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err := articleModel.Update(userID, id, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article updated"})
}

//Delete ...
func (ctrl ArticleController) Delete(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	err := articleModel.Delete(userID, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})

}
