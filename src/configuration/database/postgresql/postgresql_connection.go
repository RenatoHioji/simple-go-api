package postgresql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	USER_DATABASE_URL = "USER_DATABASE_URL"
)

func NewPostgresqlConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, os.Getenv(USER_DATABASE_URL))

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}
	return conn, nil
}
