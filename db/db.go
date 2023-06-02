package db

import (
	"os"

	"models"

	_redis "github.com/go-redis/redis/v7"
	_ "github.com/lib/pq" //import postgres
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//Init ...
func DbInit() {

	conn, err := dbSetup()
	if err != nil {
		panic(err)
	}
	db = conn
	doMigration()
}

func dbSetup() (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DATABASE_URL"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	return conn, err
}
func doMigration() {

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.ArticleModel{})
}

//RedisClient ...
var RedisClient *_redis.Client

//InitRedis ...
func InitRedis(selectDB ...int) {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

}

//GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}
