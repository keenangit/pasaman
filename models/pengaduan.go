package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

//Pengaduan ...
type Pengaduan struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primary_key:true"`
	NamaLengkap  string    `json:"namaLengkap"`
	Alamat       string    `json:"alamat"`
	NomorHP      string    `json:"nomorhp"`
	Email        string    `json:"email"`
	Pekerjaan    string    `json:"pekerjaan"`
	Tujuan       string    `json:"tujuan"`
	IsiPengaduan string    `json:"isiPengaduan"`
}

//Pengaduan ...
type PengaduanResp struct {
	ID           uuid.UUID `gorm:"primary_key:true"`
	NamaLengkap  string    `json:"namaLengkap"`
	Alamat       string    `json:"alamat"`
	NomorHP      string    `json:"nomorhp"`
	Email        string    `json:"email"`
	Pekerjaan    string    `json:"pekerjaan"`
	Tujuan       string    `json:"tujuan"`
	IsiPengaduan string    `json:"isiPengaduan"`
	CreatedAt    time.Time `json:"createdAt"`
	// UpdatedAt    time.Time `json:"updatedAt"`
}

func (u *Pengaduan) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
