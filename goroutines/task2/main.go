package main

import (
	"fmt"
	"math/rand"
)

func foo(ch chan int) {
	ch <- rand.Int()
}

func main() {
	ch := make(chan int)
	go foo(ch)
	go foo(ch)
	go foo(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
