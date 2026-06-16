package main

import (
	"CoalMine/company"
	"CoalMine/server"
	"fmt"
)

func main() {
	company := company.NewCompany()
	handlers := server.NewHTTPHandlers(company)
	server := server.NewServer(handlers)
	server.Start()

	if err := server.Start(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
