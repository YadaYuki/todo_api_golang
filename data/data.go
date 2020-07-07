package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// init
var Db *sql.DB
func init() {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	var err error
	Db, err = sql.Open("mysql", DB_USER+":"+DB_PASS+"@/todo_app?parseTime=true")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected!")
}
