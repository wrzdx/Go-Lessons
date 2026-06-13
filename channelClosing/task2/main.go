package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func main() {
	ch1 := make(chan string)
	go func() {
		n := rand.Intn(5) + 1
		for range n {
			time.Sleep(time.Duration(3 + rand.Intn(5)) * 100 * time.Millisecond)
			ch1 <- RandomString(5)
		}

		close(ch1)
	}()

	for ans := range ch1 {
		fmt.Println(ans)
	}

}
