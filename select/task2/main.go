package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		n := (1 + rand.Intn(10))*100
		time.Sleep(time.Duration(n) * time.Millisecond)
		ch1 <- n
	}()
	go func() {
		n := (1 + rand.Intn(10))*100
		time.Sleep(time.Duration(n) * time.Millisecond)
		ch2 <- n
	}()

	time.Sleep(500 * time.Millisecond)

	select {
	case number := <-ch1:
		fmt.Println("First Goroutine:", number)
	case number := <-ch2:
		fmt.Println("Second Goroutine:", number)
	default:
		fmt.Println("I was faster")
	}
}
