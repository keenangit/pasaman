package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Pendaftaran struct {
	gorm.Model
	ID uuid.UUID `gorm:"primary_key:true"`

	//Pemohon
	JenisIdentitasPemohon string `json:"jenisIdentitasPemohon"`
	NIKPemohon            string `json:"nikPemohon"`
	NamaPemohon           string `json:"namaPemohon"`
	TelpPemohon           string `json:"telpPemohon"`
	AlamatPemohon         string `json:"alamatPemohon"`
	EmailPemohon          string `json:"emailPemohon"`
	PropinsiPemohon       string `json:"propinsiPemohon"`
	KecamatanPemohon      string `json:"kecamatanPemohon"`
	KelurahanPemohon      string `json:"kelurahanPemohon"`
	KabupatenPemohon      string `json:"kabupatenPemohon"`
	KodePosPemohon        string `json:"kodePosPemohon"`

	//perusahaan
	NPWP                string `json:"npwp"`
	NoRegister          string `json:"noRegister"`
	NamaPerusahaan      string `json:"namaPerusahaan"`
	TelpPerusahaan      string `json:"telpPerusahaan"`
	AlamatPerusahaan    string `json:"alamatPerusahaan"`
	PropinsiPerusahaan  string `json:"propinsiPerusahaan"`
	KecamatanPerusahaan string `json:"kecamatanPerusahaan"`
	KelurahanPerusahaan string `json:"kelurahanPerusahaan"`
	KabupatenPerusahaan string `json:"kabupatenPerusahaan"`
	KodePosPerusahaan   string `json:"kodePosPerusahaan"`

	//Jenis Perijinan
	JenisPerizinan    string `json:"jenisPerizinan"`
	StatusPendaftaran string `json:"statusPendaftaran"`
}

type PendaftaranResp struct {
	ID uuid.UUID `gorm:"primary_key:true"`

	//Pemohon
	JenisIdentitasPemohon string `json:"jenisIdentitasPemohon"`
	NIKPemohon            string `json:"nikPemohon"`
	NamaPemohon           string `json:"namaPemohon"`
	TelpPemohon           string `json:"telpPemohon"`
	AlamatPemohon         string `json:"alamatPemohon"`
	EmailPemohon          string `json:"emailPemohon"`
	PropinsiPemohon       string `json:"propinsiPemohon"`
	KecamatanPemohon      string `json:"kecamatanPemohon"`
	KelurahanPemohon      string `json:"kelurahanPemohon"`
	KabupatenPemohon      string `json:"kabupatenPemohon"`
	KodePosPemohon        string `json:"kodePosPemohon"`

	//perusahaan
	NPWP                string `json:"npwp"`
	NoRegister          string `json:"noRegister"`
	NamaPerusahaan      string `json:"namaPerusahaan"`
	TelpPerusahaan      string `json:"telpPerusahaan"`
	AlamatPerusahaan    string `json:"alamatPerusahaan"`
	PropinsiPerusahaan  string `json:"propinsiPerusahaan"`
	KecamatanPerusahaan string `json:"kecamatanPerusahaan"`
	KelurahanPerusahaan string `json:"kelurahanPerusahaan"`
	KabupatenPerusahaan string `json:"kabupatenPerusahaan"`
	KodePosPerusahaan   string `json:"kodePosPerusahaan"`

	//Jenis Perijinan
	JenisPerizinan    string `json:"jenisPerizinan"`
	StatusPendaftaran string `json:"statusPendaftaran"`

	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
}

func (u *Pendaftaran) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	// if u.Role == "admin" {
	//   return errors.New("invalid role")
	// }
	return
}
