package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Fail to read http request body:", err)
			return
		}

		httpRequestBodyString := string(httpRequestBody)
		fmt.Println(httpRequestBodyString)
	})

	http.ListenAndServe(":8000", nil)
}
