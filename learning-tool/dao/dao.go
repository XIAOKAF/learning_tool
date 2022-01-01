package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:045226@tcp(localhost:3306)/learning_tool?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	DB = db
}
