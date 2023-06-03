package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

//PengaduanForm ...
type PengaduanForm struct{}

//CreatePengaduanForm ...
type CreatePengaduanForm struct {
	NamaLengkap  string `json:"nama_lengkap"`
	Alamat       string `json:"alamat"`
	NomorHP      string `json:"nomor_hp"`
	Email        string `json:"email"`
	Pekerjaan    string `json:"pekerjaan"`
	Tujuan       string `json:"tujuan"`
	IsiPengaduan string `json:"isi_pengaduan"`
}

//Title ...
func (f PengaduanForm) Title(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the pengaduan title"
		}
		return errMsg[0]
	case "min", "max":
		return "Title should be between 3 to 100 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

//Content ...
func (f PengaduanForm) Content(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the pengaduan content"
		}
		return errMsg[0]
	case "min", "max":
		return "Content should be between 3 to 1000 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

//Create ...
func (f PengaduanForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Title" {
				return f.Title(err.Tag())
			}
			if err.Field() == "Content" {
				return f.Content(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

//Update ...
func (f PengaduanForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Title" {
				return f.Title(err.Tag())
			}
			if err.Field() == "Content" {
				return f.Content(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
