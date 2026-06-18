package main

import (
	"encoding/json"
	"fmt"
	"io"
	"logging/logger"
	"net/http"
	"time"
)

func main() {
	logger, logFileClose, err := logger.NewLogger("DEBUG")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFileClose()
	arr := []string{}
	http.HandleFunc("/strings", func(w http.ResponseWriter, r *http.Request) {
		recievedAt := time.Now()

		switch r.Method {
		case http.MethodPost:

			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Fail to read http request body:", err)
				elapsed := time.Since(recievedAt)
				sentAt := time.Now()
				msg := fmt.Sprintf(
					"%v POST \"/strings\": Amount of strings - %d, elapsed %v, status %v, sent %v\n",
					recievedAt, len(arr), elapsed, http.StatusBadRequest, sentAt,
				)
				logger.Info(msg)
				return
			}
			arr = append(arr, string(body))
			w.WriteHeader(http.StatusCreated)
			elapsed := time.Since(recievedAt)
			sentAt := time.Now()
			msg := fmt.Sprintf(
				"%v POST \"/strings\": Amount of strings - %d, elapsed %v, status %v, sent %v\n",
				recievedAt, len(arr), elapsed, http.StatusCreated, sentAt,
			)
			logger.Info(msg)
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(arr); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				elapsed := time.Since(recievedAt)
				sentAt := time.Now()
				msg := fmt.Sprintf(
					"%v GET \"/strings\": Amount of strings - %d, elapsed %v, status %v, sent %v\n",
					recievedAt, len(arr), elapsed, http.StatusInternalServerError, sentAt,
				)
				logger.Info(msg)
				return
			}
			elapsed := time.Since(recievedAt)
			sentAt := time.Now()
			msg := fmt.Sprintf(
				"%v GET \"/strings\": Amount of strings - %d, elapsed %v, status %v, sent %v\n",
				recievedAt, len(arr), elapsed, http.StatusOK, sentAt,
			)
			logger.Info(msg)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			elapsed := time.Since(recievedAt)
			sentAt := time.Now()
			msg := fmt.Sprintf(
				"%v %v \"/strings\": Amount of strings - %d, elapsed %v, status %v, sent %v\n",
				recievedAt, r.Method, len(arr), elapsed, http.StatusMethodNotAllowed, sentAt,
			)
			logger.Info(msg)
		}
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
