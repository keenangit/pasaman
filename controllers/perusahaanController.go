package controllers

import (
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keenangit/pasaman/models"
	"github.com/keenangit/pasaman/repo"
)

type PerusahaanController struct{}

var perusahaanRepo = new(repo.PerusahaanRepo)

//Create ...
func (ctrl PerusahaanController) Create(c *gin.Context) {

	var form models.Perusahaan

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		// message := pengaduanForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Please Check"})
		return
	}

	id, err := perusahaanRepo.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Pendaftaran could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pendaftaran created", "id": id})
}

func (ctrl PerusahaanController) Export(c *gin.Context) {

	items := v3Convert.StructToSliceOfSliceWithHeader(perusahaanRepo.List())

	// Set our headers so browser will download the file
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=users.csv")
	// Create a CSV writer using our HTTP response writer as our io.Writer
	wr := csv.NewWriter(c.Writer)
	wr.Comma = ';'
	// Write all items and deal with errors
	if err := wr.WriteAll(items); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
