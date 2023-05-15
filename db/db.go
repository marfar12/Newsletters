package db

import (
	"database/sql"
	"os"

	"log"
)

type Connection struct{}

func Connect() *sql.DB {
	db, dbErr := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	return db
}
