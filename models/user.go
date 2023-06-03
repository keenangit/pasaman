package models

import (
	"errors"
	"log"
	"strings"

	uuid "github.com/google/uuid"
	"github.com/keenangit/pasaman/db"
	"github.com/keenangit/pasaman/forms"

	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID        string `db:"id, primarykey" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
	CreatedAt int64  `db:"created_at" json:"-"`
}

//UserModel ...
type UserModel struct{}

var authModel = new(AuthModel)

//Login ...
func (m UserModel) Login(form forms.LoginForm) (user User, token Token, err error) {

	err = db.GetDB().SelectOne(&user, "SELECT id, email, password, name, updated_at, created_at FROM tb_user WHERE email=LOWER(?) LIMIT 1", form.Email)

	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return user, token, err
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)
	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return user, token, err
	}

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

//Register ...
func (m UserModel) Register(form forms.RegisterForm) (user User, err error) {
	getDb := db.GetDB()

	//Check if the user exists in database
	checkUser, err := getDb.SelectInt("SELECT count(id) FROM tb_user WHERE email=LOWER(?) LIMIT 1", form.Email)
	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return user, errors.New("something went wrong, please try again later")
	}

	if checkUser > 0 {
		return user, errors.New("email already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return user, errors.New("something went wrong, please try again later")
	}

	//Create the user and return back the user ID
	id := uuid.New()
	_, err = getDb.Exec("INSERT INTO tb_user(id, email, password, name) VALUES(?, ?, ?, ?)", id, strings.ToLower(form.Email), string(hashedPassword), form.Name)
	// usr := &User{id, form.Email, string(hashedPassword), form.Name, 0, 0}
	// err = getDb.Insert(usr)

	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return user, errors.New("something went wrong, please try again later")
	}
	user.ID = id.String()
	user.Name = form.Name
	user.Email = form.Email

	return user, err
}

//One ...
func (m UserModel) One(userID string) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT id, email, name, phone FROM tb_user WHERE id=? LIMIT 1", userID)
	if err != nil {
		log.Printf("\n\n ERR: %s", err)
		return user, errors.New("something went wrong, please try again later")
	}
	return user, err
}

//GetUserByEmail ...
func GetUserByEmail(userID string) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM tb_user WHERE email=LOWER(?) LIMIT 1", userID)
	if err != nil {
		log.Printf("\n\n GetUser ERR: Email Not Found  %s", err)
		return user, errors.New("something went wrong, please try again later")
	}
	return user, err
}

func GetUserLevel(userID string) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM tb_user WHERE id=? and level=? LIMIT 1", userID, 0)
	if err != nil {
		log.Printf("\n\n GetUser ERR: Level Not Access  %s", err)
		return user, errors.New("something went wrong, please try again later")
	}
	return user, err
}
