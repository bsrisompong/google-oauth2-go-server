package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// Ensure users table exists
	ensureUsersTable()

	// Run database migrations
	// runMigrations()
}

func ensureUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		picture TEXT,
		verified_email BOOLEAN NOT NULL
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to ensure users table exists: %v", err)
	}
	log.Println("Checked/created users table successfully")
}

func runMigrations() {
	log.Println("Starting database migration...")

	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create migration driver: %v", err)
	}

	migrationPath := "file:///app/db/migrations"
	log.Printf("Migration directory path: %s", migrationPath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres", driver)

	if err != nil {
		log.Fatalf("Could not create migration instance: %v", err)
	}

	log.Println("Running migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migration: %v", err)
	}

	log.Println("Database migrations applied successfully")
}
