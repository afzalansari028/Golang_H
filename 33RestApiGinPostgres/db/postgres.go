package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	// db, err := sql.Open("mysql", "root:king@tcp(127.0.0.1:3306)/learn")
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=king dbname=postgres sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
