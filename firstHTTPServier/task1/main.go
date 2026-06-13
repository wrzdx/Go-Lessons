package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
