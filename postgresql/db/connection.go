package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var connectionString = "postgres://wrzdx:password@localhost:5432/gostudy"

func Connection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, connectionString)
}
