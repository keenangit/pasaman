package migrations

import (
	"github.com/keenangit/pasaman/db"
	"github.com/keenangit/pasaman/models"
)

func AMigrate() {
	db.GetGormDB().AutoMigrate(&models.Pendaftaran{})
	// db.GetGormDB().AutoMigrate(&models.Perusahaan{})
	db.GetGormDB().AutoMigrate(&models.Informasi{})
	db.GetGormDB().AutoMigrate(&models.Pengaduan{})
}
