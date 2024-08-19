package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}
