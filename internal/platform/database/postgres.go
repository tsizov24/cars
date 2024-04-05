package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func OpenDBConn() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), fmt.Sprintf(
		"postgres://%s:%s@db:5432/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	))
}
