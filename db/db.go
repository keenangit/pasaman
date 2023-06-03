package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_redis "github.com/go-redis/redis/v7"

	// _ "github.com/lib/pq" //import postgres
	_ "github.com/go-sql-driver/mysql"
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

//Init ...
func Init() {

	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", dataSourceName)
	// db, err := sql.Open("postgres", dataSourceName)
	checkErr(err, "db not connect")

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	// dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	// Defines the table for use by GORP
	// This is a function we will create soon.
	// dbmap.AddTableWithName(Pengaduan{}, "article").SetKeys(false, "ID")
	// err = dbmap.CreateTablesIfNotExists()
	// checkErr(err, "Create tables failed")

	return dbmap, nil
}

// type Pengaduan struct {
// 	ID           string `db:"id, primarykey" json:"id"`
// 	NamaLengkap  string `db:"nama_lengkap" json:"nama_lengkap"`
// 	Alamat       string `db:"alamat" json:"alamat"`
// 	NomorHP      string `db:"nomorhp" json:"nomorhp"`
// 	Email        string `db:"email" json:"email"`
// 	Pekerjaan    string `db:"pekerjaan" json:"pekerjaan"`
// 	Tujuan       string `db:"tujuan" json:"tujuan"`
// 	IsiPengaduan string `db:"isi_pengaduan" json:"isi_pengaduan"`
// 	UpdatedAt    int64  `db:"updated_at" json:"updated_at"`
// 	CreatedAt    int64  `db:"created_at" json:"created_at"`
// }

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
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
