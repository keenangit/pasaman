package models

import (
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
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

//Article ...
type ArticleResp struct {
	ID        string `db:"id, primarykey" json:"id"`
	UserID    string `db:"user_id" json:"-"`
	Title     string `db:"title" json:"title"`
	Content   string `db:"content" json:"content"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
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

	operation, err := db.GetDB().Exec("INSERT INTO tb_article(id, user_id, title, content, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)", id, userID, form.Title, form.Content, now.Unix(), now.Unix())
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
	// err = db.GetDB().SelectOne(&article, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 AND a.id=$2 LIMIT 1", userID, id)
	var art Article
	err := db.GetDB().SelectOne(&art, "SELECT title FROM tb_article where user_id=? and id=?", userID, id)
	checkErr(err, "Select failed")

	return art, err
}

//All ...
func (m ArticleModel) GetAll(userID string) ([]ArticleResp, error) {
	var art []ArticleResp
	_, err := db.GetDB().Select(&art, "SELECT * FROM tb_article where user_id=?", userID)
	checkErr(err, "Select failed")

	return art, err
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Printf(msg, err)
	}
}

//Update ...
func (m ArticleModel) Update(userID string, id int64, form forms.CreateArticleForm) (err error) {
	//METHOD 1
	//Check the article by ID using this way
	// _, err = m.One(userID, id)
	// if err != nil {
	// 	return err
	// }

	operation, err := db.GetDB().Exec("UPDATE public.article SET title=$2, content=$3 WHERE id=$1", id, form.Title, form.Content)
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

	operation, err := db.GetDB().Exec("DELETE FROM public.article WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}
