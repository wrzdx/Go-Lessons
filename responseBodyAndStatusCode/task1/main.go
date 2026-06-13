package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		number, err := strconv.Atoi(string(body))
		if err != nil {
			fmt.Println("Err:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch number {
		case 200:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
		}
	})

	if err:= http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Err:", err)
	}
}
