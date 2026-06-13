package main

import (
	"fmt"
	"net/http"
)

func Dog(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Я собака и я говорю 'Гав'"))
	if err != nil {
		fmt.Println("Dog handler err:", err)
		return
	}

	fmt.Println("Dog handler")
}


func Cat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Я кот и я говорю 'Мяу'"))
	if err != nil {
		fmt.Println("Cat handler err:", err)
		return
	}

	fmt.Println("Cat handler")
}


func Cow(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Я корова и я говорю 'Муу'"))
	if err != nil {
		fmt.Println("Cow handler err:", err)
		return
	}

	fmt.Println("Cow handler")
}

func main() {
	http.HandleFunc("/dog", Dog)
	http.HandleFunc("/cat", Cat)
	http.HandleFunc("/cow", Cow)

	if err:=http.ListenAndServe(":8000",nil); err!= nil {
		fmt.Println(err)
	}
}
