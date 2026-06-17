package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		title VARCHAR(200) NOT NULL,
		author VARCHAR(200) NOT NULL,
		review VARCHAR(1000),
		year INTEGER,
		is_read BOOLEAN NOT NULL,
		added_at TIMESTAMP NOT NULL,
		read_at TIMESTAMP
	)
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}

func CreateBook(ctx context.Context, conn *pgx.Conn, b BookModel) error {
	sqlQuery := `
	INSERT INTO books 
	(title, author, review, year, is_read, added_at, read_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		b.Title,
		b.Author,
		b.Review,
		b.Year,
		b.IsRead,
		b.AddedAt,
		b.ReadAt,
	)

	return err
}

func GetBooks(ctx context.Context, conn *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
	SELECT * FROM books;
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []BookModel{}
	for rows.Next() {
		book := BookModel{}

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.Year,
			&book.IsRead,
			&book.AddedAt,
			&book.ReadAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func UpdateBook(ctx context.Context, conn *pgx.Conn, b BookModel) error {
	sqlQuery := `
	UPDATE books SET 
	title=$1, author=$2, review=$3, year=$4, is_read=$5, added_at=$6, read_at=$7
	WHERE id=$8; 
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		b.Title,
		b.Author,
		b.Review,
		b.Year,
		b.IsRead,
		b.AddedAt,
		b.ReadAt,
		b.ID,
	)

	return err
}

func DeleteBooks(ctx context.Context, conn *pgx.Conn, ids []int) error {
	sqlQuery := `
	DELETE FROM books WHERE id=ANY($1);
	`
	_, err := conn.Exec(ctx, sqlQuery, ids)

	return err
}

func GetPage(ctx context.Context, conn *pgx.Conn, limit int, offset int) ([]BookModel, error) {
	sqlQuery := `
	SELECT * FROM books LIMIT $1 OFFSET $2;
	`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []BookModel{}
	for rows.Next() {
		book := BookModel{}

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.Year,
			&book.IsRead,
			&book.AddedAt,
			&book.ReadAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
func ListPages(ctx context.Context, conn *pgx.Conn, n int) {
	for i := 0; ; i++ {
		page, err := GetPage(ctx, conn, n, i*n)
		if err != nil {
			panic(err)
		}
		if len(page) == 0 {
			break
		}
		fmt.Printf("Page %d: ", i+1)
		pp.Println(page)
	}
}
