package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	uuid "github.com/google/uuid"
	"github.com/keenangit/pasaman/db"
	"github.com/keenangit/pasaman/forms"
)

//Article ...
type Article struct {
	ID        string `db:"id, primarykey" json:"id"`
	UserID    string `db:"user_id" json:"-"`
	Title     string `db:"title" json:"title"`
	Content   string `db:"content" json:"content"`
	UrlPhoto  string `db:"url_photo" json:"url_photo"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	DeletedAt int64  `db:"deleted_at" json:"-"`
}

//Article ...
type ArticleResp struct {
	ID string `db:"id, primarykey" json:"id"`
	// UserID    string         `db:"user_id" json:"-"`
	Title    string         `db:"title" json:"title"`
	Content  string         `db:"content" json:"content"`
	UrlPhoto sql.NullString `db:"url_photo" json:"url_photo"`
	// UpdatedAt int64          `db:"updated_at" json:"updated_at"`
	// CreatedAt int64          `db:"created_at" json:"created_at"`
	// DeletedAt sql.NullString `db:"deleted_at" json:"-"`
}

//ArticleModel ...
type ArticleModel struct{}

//Create ...
func (m ArticleModel) Create(userID string, form forms.CreateArticleForm) (articleID string, err error) {
	now := time.Now()
	id := uuid.New()
	articleID = id.String()

	// data := &Article{
	// 	ID:        uuid.New().String(),
	// 	UserID:    userID,
	// 	Title:     form.Title,
	// 	Content:   form.Content,
	// 	CreatedAt: now.Unix(),
	// 	UpdatedAt: now.Unix(),
	// }
	// err = db.GetDB().Insert(data)

	operation, err := db.GetDB().Exec("INSERT INTO tb_article(id, user_id, title, content, url_photo, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)", id, userID, form.Title, form.Content, form.UrlPhoto, now.Unix(), now.Unix())
	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return articleID, err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return articleID, errors.New("no records insert")
	}
	return articleID, err
}

//One ...
func (m ArticleModel) One(userID string, id string) (Article, error) {
	// err = db.GetDB().SelectOne(&article, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM article a LEFT JOIN user u ON a.user_id = u.id WHERE a.user_id=$1 AND a.id=$2 LIMIT 1", userID, id)
	var art Article
	err := db.GetDB().SelectOne(&art, "SELECT id, title, content, url_photo FROM tb_article where id=?", id)
	checkErr(err, "Select failed")

	return art, err
}

//All ...
func (m ArticleModel) GetAll(userID string) ([]ArticleResp, error) {
	var art []ArticleResp
	_, err := db.GetDB().Select(&art, "SELECT id, title, content, url_photo FROM tb_article where deleted_at is null order by created_at desc")
	checkErr(err, "Select failed")

	return art, err
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Printf(msg, err)
	}
}

//Update ...
func (m ArticleModel) Update(userID string, id string, form forms.CreateArticleForm) (err error) {

	operation, err := db.GetDB().Exec("UPDATE tb_article SET title=?, content=?, url_photo=? WHERE id=?", form.Title, form.Content, form.UrlPhoto, id)
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
func (m ArticleModel) Delete(userID, id string) (err error) {
	now := time.Now()
	// operation, err := db.GetDB().Exec("DELETE FROM tb_article WHERE id=$1", id)
	operation, err := db.GetDB().Exec("UPDATE tb_article SET deleted_at=? WHERE id=?", now.Unix(), id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}
