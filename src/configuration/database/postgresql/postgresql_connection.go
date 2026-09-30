package postgresql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	USER_DATABASE_URL = "USER_DATABASE_URL"
)

func NewPostgresqlConnection(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, os.Getenv(USER_DATABASE_URL))

	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	return pool, nil
}
