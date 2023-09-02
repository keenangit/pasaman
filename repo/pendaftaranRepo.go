package repo

import (
	"log"

	"github.com/keenangit/pasaman/db"
	"github.com/keenangit/pasaman/models"
)

type PendaftaranRepo struct{}

//Create ...
func (m PendaftaranRepo) Create(form models.Pendaftaran) (PendaftaranID string, err error) {

	result := db.GetGormDB().Create(&form)

	return form.ID.String(), result.Error
}

//Create ...
func (m PendaftaranRepo) List() (mP []models.PendaftaranResp) {

	result := db.GetGormDB().Table("pendaftarans").Find(&mP)

	log.Printf("\n\n ERR: %s", result.Error)

	return mP
}
