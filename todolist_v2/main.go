package main

import (
	"context"
	"fmt"
	"os"
	"restapi/internal/repository"
	"restapi/internal/service"
	"restapi/internal/transport"

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
	taskService := service.NewTaskService(repo)
	handlers := transport.NewHTTPHandlers(taskService)
	server := transport.NewHTTPServer(handlers)
	fmt.Println("Server successfuly started!")
	if err := server.StartServer(); err != nil {
		fmt.Println(err)
	}
}
