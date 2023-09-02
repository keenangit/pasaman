package repo

import (
	"log"

	"github.com/keenangit/pasaman/db"
	"github.com/keenangit/pasaman/models"
)

type InformasiRepo struct{}

//Create ...
func (m InformasiRepo) Create(form models.Informasi) (InformasiID string, err error) {

	result := db.GetGormDB().Create(&form)

	return form.ID.String(), result.Error
}

//Create ...
func (m InformasiRepo) List() (mP []models.InformasiResp) {

	result := db.GetGormDB().Table("informasis").Find(&mP)

	log.Printf("\n\n ERR: %s", result.Error)

	return mP
}
