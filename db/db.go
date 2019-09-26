package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init(dbName string) {
	var err error

	db, _ = sql.Open("mysql", "root:password@tcp(db:3306)/"+dbName+"?collation=utf8mb4_general_ci")
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Print("Connection with database successful")
	}
}

func GetDB() *sql.DB {
	return db
}
