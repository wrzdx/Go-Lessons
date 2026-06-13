package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		for k,v := range r.Header {
			fmt.Printf("%s: %q\n", k,v)
		}

		if name, ok := r.Header["Name"]; ok {
			fmt.Println("Hello,", name[0])
		}
	})

	if err:=http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}