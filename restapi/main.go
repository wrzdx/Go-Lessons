package main

import (
	"fmt"
	"library/library"
	"library/server"
)

func main() {
	library := library.NewLibrary()
	handlers := server.NewHTTPHandlers(library)
	server := server.NewServer(handlers)
	server.Start()

	if err := server.Start(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
