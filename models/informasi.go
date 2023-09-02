package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Informasi struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primary_key:true"`
	Nama             string    `json:"nama"`
	Alamat           string    `json:"alamat"`
	Pekerjaan        string    `json:"pekerjaan"`
	Telp             string    `json:"telp"`
	Email            string    `json:"email"`
	NomorIdentitas   string    `json:"nomorIdentitas"`
	RincianInformasi string    `json:"rincianInformasi"`
	TujuanInformasi  string    `json:"tujuanInformasi"`
	CaraMemperoleh   string    `json:"caraMemperoleh"`
	BentukSalinan    string    `json:"bentukSalinan"`
	CaraMendapatkan  string    `json:"caraMendapatkan"`
}

type InformasiResp struct {
	ID               uuid.UUID `gorm:"primary_key:true"`
	Nama             string    `json:"nama"`
	Alamat           string    `json:"alamat"`
	Pekerjaan        string    `json:"pekerjaan"`
	Telp             string    `json:"telp"`
	Email            string    `json:"email"`
	NomorIdentitas   string    `json:"nomorIdentitas"`
	RincianInformasi string    `json:"rincianInformasi"`
	TujuanInformasi  string    `json:"tujuanInformasi"`
	CaraMemperoleh   string    `json:"caraMemperoleh"`
	BentukSalinan    string    `json:"bentukSalinan"`
	CaraMendapatkan  string    `json:"caraMendapatkan"`
	CreatedAt        time.Time `json:"createdAt"`
}

func (u *Informasi) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
