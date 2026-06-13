package main

import "fmt"

func main() {
	var ch chan int = make(chan int)

	close(ch)
	ch <- 1
	fmt.Println(<- ch)
}