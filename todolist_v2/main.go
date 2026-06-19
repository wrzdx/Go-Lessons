package main

import (
	"context"
	"fmt"
	"os"
	"restapi/internal/core"
	"restapi/internal/repository"
	"restapi/internal/service"
	"restapi/internal/transport"

	"github.com/jackc/pgx/v5"
)

func main() {
	logger, logFileClose, err := core.NewLogger("INFO")
	if err != nil {
		panic(err)
	}
	defer logFileClose()
	logger.Info("Logger initialized successfully, file rotation simulated")
	ctx := context.Background()
	connString := os.Getenv("LOCAL_DATABASE_URL")
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Println(err)
		return
	}
	repo := repository.NewPostgres(conn)
	taskService := service.NewTaskService(repo)
	handlers := transport.NewHTTPHandlers(taskService,logger)
	server := transport.NewHTTPServer(handlers)
	if err := server.StartServer(); err != nil {
		fmt.Println(err)
	}
}
