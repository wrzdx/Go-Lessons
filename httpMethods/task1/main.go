package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get", func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Println(errors.New("Wrong method: " + r.Method))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("Right method:", r.Method)

	})

	http.HandleFunc("/post", func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println(errors.New("Wrong method: " + r.Method))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("Right method:", r.Method)
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error:", err)
	}
}