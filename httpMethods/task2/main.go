package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Message struct {
	ID      string
	Content string
}

func main() {
	mutex := sync.Mutex{}
	msgs := []Message{}
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println(errors.New("Wrong method: " + r.Method))
			w.WriteHeader(http.StatusMethodNotAllowed)
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

		body := string(httpRequestBody)
		newMsg := Message{
			ID:      strconv.FormatUint(rand.Uint64(), 16),
			Content: body,
		}
		mutex.Lock()
		msgs = append(msgs, newMsg)
		mutex.Unlock()
		res := fmt.Sprintf("Message created: {ID: %s, Content: %s}", newMsg.ID, newMsg.Content)
		w.WriteHeader(http.StatusCreated)
		if _, err := io.WriteString(w, res); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
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
			res := fmt.Sprint(msgs)
			mutex.Unlock()
			if _, err := io.WriteString(w, res); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
			return
		}
		mutex.Lock()
		found := false
		for _, msg := range msgs {
			if msg.ID == id {
				res := fmt.Sprintf("Message: {ID: %s, Content: %s}", msg.ID, msg.Content)
				if _, err := io.WriteString(w, res); err != nil {
					res = err.Error()
				}
				fmt.Println(res)
				found = true
				break
			}
		}
		mutex.Unlock()
		if !found {
			w.WriteHeader(http.StatusNotFound)
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

		body := string(httpRequestBody)
		mutex.Lock()

		found := false
		for i, msg := range msgs {
			if msg.ID == body {
				msgs = append(msgs[:i], msgs[i+1:]...)
				res := fmt.Sprintf("Message deleted: {ID: %s, Content: %s}", msg.ID, msg.Content)
				if _, err := io.WriteString(w, res); err != nil {
					res = err.Error()
				}
				fmt.Println(res)
				found = true
				break
			}
		}
		mutex.Unlock()
		if !found {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(":8000", nil)
}
