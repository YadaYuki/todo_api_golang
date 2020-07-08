package data

import (
	"database/sql"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

// init
var Db *sql.DB
func init() {
	err := godotenv.Load(".env.development")
	// https://teratail.com/questions/106823 envファイルの取り扱い!
	// testの際おそらくエラーになるので注意!
	if err != nil{
		log.Fatal(err)
		return
	}
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	log.Println(DB_USER+","+DB_PASS)
	Db, err = sql.Open("mysql", DB_USER+":"+DB_PASS+"@/todo_app?parseTime=true")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Println("connected!")
}
