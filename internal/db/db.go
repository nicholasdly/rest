package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nicholasdly/rest/internal/config"
)

func NewPool(ctx context.Context, config config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, config.DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("Unable to create database pool: %w", err)
	}
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Unable to ping database: %w", err)
	}

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			username TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);

		CREATE OR REPLACE FUNCTION set_updated_at()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.updated_at = NOW();
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE OR REPLACE TRIGGER users_set_updated_at
		BEFORE UPDATE ON users
		FOR EACH ROW
		EXECUTE FUNCTION set_updated_at();
	`)

	if err != nil {
		return nil, fmt.Errorf("Failed to run database migrations: %w", err)
	}

	return pool, nil
}
