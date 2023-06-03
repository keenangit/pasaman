package controllers

import (
	"strconv"

	"github.com/keenangit/pasaman/forms"
	"github.com/keenangit/pasaman/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

//PengaduanController ...
type PengaduanController struct{}

var pengaduanModel = new(models.PengaduanModel)
var pengaduanForm = new(forms.PengaduanForm)

//Create ...
func (ctrl PengaduanController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreatePengaduanForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := pengaduanForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := pengaduanModel.Create(userID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Pengaduan could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan created", "id": id})
}

//All ...
func (ctrl PengaduanController) All(c *gin.Context) {
	userID := getUserID(c)

	results, err := pengaduanModel.GetAll(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get pengaduans"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

//One ...
func (ctrl PengaduanController) One(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	data, err := pengaduanModel.One(userID, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Pengaduan not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

//Update ...
func (ctrl PengaduanController) Update(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var form forms.CreatePengaduanForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := pengaduanForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = pengaduanModel.Update(userID, getID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Pengaduan could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan updated"})
}

//Delete ...
func (ctrl PengaduanController) Delete(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	err := pengaduanModel.Delete(userID, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Pengaduan could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengaduan deleted"})

}
