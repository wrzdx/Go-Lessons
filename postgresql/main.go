package main

import (
	"context"
	"postgres/db"
)

func main() {
	ctx := context.Background()
	conn, err := db.Connection(ctx)
	if err != nil {
		panic(err)
	}

	db.ListPages(ctx, conn, 10)

}
