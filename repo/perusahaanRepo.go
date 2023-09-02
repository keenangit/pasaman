package repo

import (
	"log"

	"github.com/keenangit/pasaman/db"
	"github.com/keenangit/pasaman/models"
)

type PerusahaanRepo struct{}

//Create ...
func (m PerusahaanRepo) Create(form models.Perusahaan) (PerusahaanID string, err error) {

	result := db.GetGormDB().Create(&form)

	return form.ID.String(), result.Error
}

//Create ...
func (m PerusahaanRepo) List() (mP []models.PerusahaanResp) {

	result := db.GetGormDB().Table("perusahaans").Find(&mP)

	log.Printf("\n\n ERR: %s", result.Error)

	return mP
}
