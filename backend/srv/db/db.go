package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Office-Stapler/Palplan/backend/srv/config"
	"github.com/jackc/pgx/v5"
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

// Exec executes a SQL command without returning rows.
func (db *DB) Exec(ctx context.Context, sql string, arguments ...any) error {
	if db.Pool == nil {
		return fmt.Errorf("database pool is nil")
	}
	_, err := db.Pool.Exec(ctx, sql, arguments...)
	return err
}

// Query executes a SQL query that returns rows.
func (db *DB) Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error) {
	if db.Pool == nil {
		return nil, fmt.Errorf("database pool is nil")
	}
	return db.Pool.Query(ctx, sql, arguments...)
}

// QueryRow executes a SQL query that returns at most one row.
func (db *DB) QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row {
	if db.Pool == nil {
		log.Println("database pool is nil") //Log instead of error, since this function returns a row.
		return nil
	}
	return db.Pool.QueryRow(ctx, sql, arguments...)
}
