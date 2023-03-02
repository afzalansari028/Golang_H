package ConnectionProvider

import (
	"database/sql"
	"log"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "king"
	DB_NAME     = "library"
)

// DB set up
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:king@tcp(127.0.0.1:3306)/learn")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
