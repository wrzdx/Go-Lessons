package main

import (
	"fmt"
	"time"
)

func foo(n int) {
	for i := range 5 {
		fmt.Printf("I'm goroutine %d, printing %d times\n", n, i + 1)
		time.Sleep(time.Second)
	}
}

func main() {
	go foo(1)
	go foo(2)
	go foo(3)

	time.Sleep(5 * time.Second + 100* time.Millisecond)
}