package postgresql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitPostgresql() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		panic(err)
	}

	return conn
}
