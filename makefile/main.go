package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		fmt.Println(errors.New("DATABASE_URL is not set"))
		return
	}
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		fullName VARCHAR(200) NOT NULL,
		phoneNumber VARCHAR(200)
	);
	`
	_, err = conn.Exec(ctx, sqlQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	newUser := os.Getenv("NEW_USER")
	if newUser == "YES" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Fullname: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println(scanner.Err())
			return
		}
		fullname := scanner.Text()
		fmt.Print("Phone Number: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println(scanner.Err())
			return
		}
		phoneNumber := scanner.Text()
		sqlQuery = `
		INSERT INTO users (fullName, phoneNumber) VALUES ($1, $2)
		`

		_, err = conn.Exec(ctx, sqlQuery, fullname, phoneNumber)
		if err != nil {
			fmt.Println(err)
		}
	} else if newUser == "NO" {
		sqlQuery = `
		SELECT * FROM users;
		`
		rows, err := conn.Query(ctx, sqlQuery)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user struct {
				ID          int
				FullName    string
				PhoneNumber string
			}

			if err := rows.Scan(&user.ID, &user.FullName, &user.PhoneNumber); err != nil {
				fmt.Println(err)
				return
			}

			pp.Println(user)
		}
	} else {
		fmt.Println("Invalid NEW_USER variable")
	}
}
