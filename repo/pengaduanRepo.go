package repo

import (
	"log"

	"github.com/keenangit/pasaman/db"
	"github.com/keenangit/pasaman/models"
)

type PengaduanRepo struct{}

//Create ...
func (m PengaduanRepo) Create(form models.Pengaduan) (PengaduanID string, err error) {

	result := db.GetGormDB().Create(&form)

	return form.ID.String(), result.Error
}

//Create ...
func (m PengaduanRepo) List() (mP []models.PengaduanResp) {

	result := db.GetGormDB().Table("pengaduans").Find(&mP)

	log.Printf("\n\n ERR: %s", result.Error)

	return mP
}
