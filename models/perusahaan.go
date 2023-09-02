package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Perusahaan struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primary_key:true"`
	NPWP       string    `json:"npwp"`
	NoRegister string    `json:"noRegister"`
	Nama       string    `json:"nama"`
	Telp       string    `json:"telp"`
	Alamat     string    `json:"alamat"`
	Propinsi   string    `json:"propinsi"`
	Kecamatan  string    `json:"kecamatan"`
	Kelurahan  string    `json:"kelurahan"`
	Kabupaten  string    `json:"kabupaten"`
	KodePos    string    `json:"kodePos"`
}

type PerusahaanResp struct {
	ID         uuid.UUID `gorm:"primary_key:true"`
	NPWP       string    `json:"npwp"`
	NoRegister string    `json:"noRegister"`
	Nama       string    `json:"nama"`
	Telp       string    `json:"telp"`
	Alamat     string    `json:"alamat"`
	Propinsi   string    `json:"propinsi"`
	Kecamatan  string    `json:"kecamatan"`
	Kelurahan  string    `json:"kelurahan"`
	Kabupaten  string    `json:"kabupaten"`
	KodePos    string    `json:"kodePos"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (u *Perusahaan) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
