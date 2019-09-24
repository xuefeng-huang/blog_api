package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	var err error

	db, err = sql.Open("mysql", "root:password@tcp(mysql:3306)/blog")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Print("Connection with database successful")
	}
}

func GetDB() *sql.DB {
	return db
}
