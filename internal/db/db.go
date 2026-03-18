package db

import (
	"database/sql"
	"log/slog"

	_ "modernc.org/sqlite"
)

func New(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		slog.Error("failed to open sqlite db", "err", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		slog.Error("failed to ping sqlite db", "err", err)
		return nil, err
	}

	db.SetMaxOpenConns(1)

	if err = migrate(db); err != nil {
		slog.Error("failed to run migration queries", "err", err)
		return nil, err
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		return err
	}

	return nil
}
