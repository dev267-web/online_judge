package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres wraps the pgx connection pool
type Postgres struct {
	Pool *pgxpool.Pool
}

// NewPostgres creates a new Postgres connection
func NewPostgres(ctx context.Context, dsn string) (*Postgres, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Postgres{Pool: pool}, nil
}
