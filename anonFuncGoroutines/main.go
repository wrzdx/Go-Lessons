package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("I'm supposed to me first!")
	}()
	go func() {
		fmt.Println("I'm supposed to me second!")
	}()
	go func() {
		fmt.Println("I'm supposed to me third!")
	}()

	time.Sleep(100 * time.Millisecond)
}
