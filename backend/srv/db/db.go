package db

import (
	"context"
	"fmt"

	"github.com/Office-Stapler/Palplan/backend/srv/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is a wrapper around pgxpool.Pool.
type DB struct {
	Pool *pgxpool.Pool
}

// NewDB creates a new DB instance and establishes a connection to the PostgreSQL database.
func NewDB(ctx context.Context, config *config.Config) (*DB, error) {

	if config == nil || config.DBSource == "" {
		return nil, fmt.Errorf("Empty config, please setup correctly")
	}

	pool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &DB{Pool: pool}, nil
}

// Close closes the connection pool.
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

// GetPool returns the pgxpool.Pool.
func (db *DB) GetPool() *pgxpool.Pool {
	return db.Pool
}
