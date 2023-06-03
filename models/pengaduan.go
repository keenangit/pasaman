package models

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	uuid "github.com/google/uuid"
	"github.com/keenangit/pasaman/db"
	"github.com/keenangit/pasaman/forms"
)

//Pengaduan ...
type Pengaduan struct {
	ID           string `db:"id, primarykey" json:"id"`
	NamaLengkap  string `db:"nama_lengkap" json:"nama_lengkap"`
	Alamat       string `db:"alamat" json:"alamat"`
	NomorHP      string `db:"nomorhp" json:"nomorhp"`
	Email        string `db:"email" json:"email"`
	Pekerjaan    string `db:"pekerjaan" json:"pekerjaan"`
	Tujuan       string `db:"tujuan" json:"tujuan"`
	IsiPengaduan string `db:"isi_pengaduan" json:"isi_pengaduan"`
	UpdatedAt    int64  `db:"updated_at" json:"updated_at"`
	CreatedAt    int64  `db:"created_at" json:"created_at"`
}

//Pengaduan ...
type PengaduanResp struct {
	ID        string `db:"id, primarykey" json:"id"`
	UserID    string `db:"user_id" json:"-"`
	Title     string `db:"title" json:"title"`
	Content   string `db:"content" json:"content"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

//PengaduanModel ...
type PengaduanModel struct{}

//Create ...
func (m PengaduanModel) Create(userID string, form forms.CreatePengaduanForm) (pengaduanID string, err error) {
	now := time.Now()
	id := uuid.New()
	pengaduanID = id.String()

	e, err := json.Marshal(form)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(e))

	operation, err := db.GetDB().Exec("INSERT INTO tb_pengaduan(id, nama_lengkap, alamat, nomorhp, email, pekerjaan, tujuan, isi_pengaduan, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ? ,?, ?, ?)",
		id, form.NamaLengkap, form.Alamat, form.NomorHP, form.Email, form.Pekerjaan, form.Tujuan, form.IsiPengaduan, now.Unix(), now.Unix())
	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return pengaduanID, err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return pengaduanID, errors.New("no records insert")
	}
	return pengaduanID, err
}

//One ...
func (m PengaduanModel) One(userID string, id string) (Pengaduan, error) {
	// err = db.GetDB().SelectOne(&pengaduan, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.pengaduan a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 AND a.id=$2 LIMIT 1", userID, id)
	var art Pengaduan
	err := db.GetDB().SelectOne(&art, "SELECT * FROM tb_pengaduan where user_id=? and id=?", userID, id)
	checkErr(err, "Select failed")

	return art, err
}

//All ...
func (m PengaduanModel) GetAll(userID string) ([]PengaduanResp, error) {
	var art []PengaduanResp
	_, err := db.GetDB().Select(&art, "SELECT * FROM tb_pengaduan where user_id=?", userID)
	checkErr(err, "Select failed")

	return art, err
}

//Update ...
func (m PengaduanModel) Update(userID string, id int64, form forms.CreatePengaduanForm) (err error) {
	//METHOD 1
	//Check the pengaduan by ID using this way
	// _, err = m.One(userID, id)
	// if err != nil {
	// 	return err
	// }

	operation, err := db.GetDB().Exec("UPDATE pengaduan SET title=$2, content=$3 WHERE id=$1", id, form.NamaLengkap, form.NomorHP)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

//Delete ...
func (m PengaduanModel) Delete(userID, id string) (err error) {

	operation, err := db.GetDB().Exec("DELETE FROM pengaduan WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}
