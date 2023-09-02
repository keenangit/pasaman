package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbGorm *gorm.DB

func InitGorm() {
	var err error

	// dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dbGorm, err = gorm.Open(mysql.Open(dbinfo), &gorm.Config{})
	if err != nil {
		panic("failed Gorm to connect database")
	}

	// dbGorm.AutoMigrate(&Product{})

}

func GetGormDB() *gorm.DB {
	return dbGorm
}