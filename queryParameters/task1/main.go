package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("param:", r.URL.Query().Get("param"))
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error:", err)
	}
}