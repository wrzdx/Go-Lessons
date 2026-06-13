package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan string), make(chan float64)

	go func() {
		for {
			ch1 <- rand.Int()
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Hello!"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			ch3 <- rand.Float64()
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		select {
		case number := <-ch1:
			fmt.Println(number)
		case str := <-ch2:
			fmt.Println(str)
		case float := <-ch3:
			fmt.Println(float)
		}
	}
}
