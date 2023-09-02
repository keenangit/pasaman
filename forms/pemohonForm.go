package forms

type CreatePemohonForm struct {
	JenisIdentitas string `json:"jenis_identitas"`
	NIK            string `json:"nik"`
	Nama           string `json:"nama"`
	Telp           string `json:"telp"`
	Alamat         string `json:"alamnat"`
	Email          string `json:"email"`
}
