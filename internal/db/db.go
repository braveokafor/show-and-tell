package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// Connect establishes a connection to the PostgreSQL database
func Connect() (*sql.DB, error) {
	// Use DATABASE_URL environment variable
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/showtell?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
