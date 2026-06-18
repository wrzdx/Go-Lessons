package main

import (
	"context"
	"docker/worker"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			newWorker := worker.Worker{}
			if err := json.NewDecoder(r.Body).Decode(&newWorker); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println(err)
				return
			}
			query := "INSERT INTO workers (fullName, position) VALUES ($1, $2);"
			if _, err := conn.Exec(ctx, query, newWorker.FullName, newWorker.Position); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
			w.WriteHeader(http.StatusCreated)
		case http.MethodGet:
			workers := []worker.Worker{}
			rows, err := conn.Query(ctx, "SELECT * FROM workers;")
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			for rows.Next() {
				worker := worker.Worker{}
				err := rows.Scan(&worker.ID, &worker.FullName, &worker.Position)
				if err != nil {
					fmt.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				workers = append(workers, worker)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(workers)
		case http.MethodDelete:
			var deleteWorker struct {
				ID int
			}
			if err := json.NewDecoder(r.Body).Decode(&deleteWorker); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println(err)
				return
			}
			query := `
			DELETE FROM workers WHERE id=$1;
			`
			if _, err := conn.Exec(ctx, query, deleteWorker.ID); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	})

	fmt.Println("Server successfully started!")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
