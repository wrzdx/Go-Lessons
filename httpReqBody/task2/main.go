package main

import (
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"
)

type Message struct {
	ID      string
	content string
}

func main() {
	msgs := []Message{}
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Fail to read http request body:", err)
			return
		}

		httpRequestBodyString := string(httpRequestBody)
		msgs = append(msgs, Message{
			strconv.FormatUint(rand.Uint64(), 16),
			httpRequestBodyString,
		})

		fmt.Println(msgs)
	})

	http.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Fail to read http request body:", err)
			return
		}

		httpRequestBodyString := string(httpRequestBody)

		for i, msg := range msgs {
			if msg.ID == httpRequestBodyString {
				msgs = append(msgs[:i], msgs[i+1:]...)
				break
			}
		}

		fmt.Println(msgs)
	})

	http.ListenAndServe(":8000", nil)
}
