package main

import (
	"fmt"
	"restapi/internal/service"
	"restapi/internal/transport"
)

func main() {
	todoList := service.NewTaskService()
	httpHandlers := transport.NewHTTPHandlers(todoList)
	httpServer := transport.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}