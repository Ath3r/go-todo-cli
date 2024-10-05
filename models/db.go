package models

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./db.sqlite3")
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil

}

func CloseDB(db *sql.DB) error {
	return db.Close()
}