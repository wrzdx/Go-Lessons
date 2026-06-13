package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k0kubun/pp"
)

type User struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Age       int     `json:"age"`
	IsMarried bool    `json:"isMarried"`
	Height    float64 `json:"height"`
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pp.Println(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:      "Evgenii",
		Age:       19,
		Height:    175,
		IsMarried: false,
		Address:   "Meta Residence",
	}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/read", ReadUser)
	http.HandleFunc("/get", GetUser)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error:", err)
	}
}
