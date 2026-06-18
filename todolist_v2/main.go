package main

import (
	"context"
	"fmt"
	"os"
	"restapi/internal/repository"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	connString := os.Getenv("LOCAL_DATABASE_URL")
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Println(err)
		return
	}
	repo := repository.NewPostgres(conn)
	err = repo.Delete(ctx, "empty")
	if err != nil {
		fmt.Println(err)
		return
	}
}
