package db

import (
	"database/sql"
	"fmt"
	"log"
	"newsletter/config"
)

type Connection struct{}

func Connect(cfg config.Config) *sql.DB {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName)
	db, dbErr := sql.Open("postgres", connectionString)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	return db
}
