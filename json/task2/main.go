package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Message struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Postcode int    `json:"postcode"`
	IsUrgent bool   `json:"isUrgent"`
}

type MessageResponse struct {
	ID string `json:"id"`
	Message
}

func main() {
	mutex := sync.Mutex{}
	msgs := map[string]Message{}
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println(errors.New("Wrong method: " + r.Method))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var newMsg Message
		if err := json.NewDecoder(r.Body).Decode(&newMsg); err != nil {
			fmt.Println("err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		id := strconv.FormatUint(rand.Uint64(), 16)
		mutex.Lock()
		msgs[id] = newMsg
		mutex.Unlock()
		b, err := json.Marshal(MessageResponse{ID: id, Message: newMsg})
		if err != nil {
			fmt.Println("err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(b); err != nil {
			fmt.Println("err:", err)
			return
		}
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Println(errors.New("Wrong method: " + r.Method))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			res := "Fail to read http request body: " + err.Error()
			fmt.Println(res)
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := io.WriteString(w, res); err != nil {
				fmt.Println(err)
			}
			return
		}

		id := strings.TrimSpace(string(body))
		if len(id) == 0 {
			mutex.Lock()

			res := make([]MessageResponse, 0, len(msgs))

			for id, msg := range msgs {
				res = append(res, MessageResponse{
					ID:      id,
					Message: msg,
				})
			}
			b, err := json.Marshal(res)
			mutex.Unlock()
			if err != nil {
				fmt.Println("err:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if _, err := w.Write(b); err != nil {
				fmt.Println("err:", err)
			}
			return
		}
		mutex.Lock()
		msg, ok := msgs[id]
		mutex.Unlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		b, err := json.Marshal(MessageResponse{ID: id, Message: msg})
		if err != nil {
			fmt.Println("err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(b); err != nil {
			fmt.Println("err:", err)
		}
	})

	http.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			res := "Fail to read http request body: " + err.Error()
			fmt.Println(res)
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := io.WriteString(w, res); err != nil {
				fmt.Println(err)
			}
			return
		}

		id := string(httpRequestBody)
		mutex.Lock()
		deleted, ok := msgs[id]
		delete(msgs, id)
		mutex.Unlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		b, err := json.Marshal(MessageResponse{ID: id, Message: deleted})
		if err != nil {
			fmt.Println("err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(b); err != nil {
			fmt.Println("err:", err)
			return
		}
	})

	http.ListenAndServe(":8000", nil)
}
