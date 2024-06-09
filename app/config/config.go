package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(localhost:3306)/goapi"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %s", err.Error())
	} else {
		fmt.Println("Test connection is successful")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
}

func GetDB() *sql.DB {
	return DB
}
