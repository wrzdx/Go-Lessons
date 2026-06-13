package main

import "fmt"


func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("There was panic:", p)
		}
	}()
	panic(342)
}
