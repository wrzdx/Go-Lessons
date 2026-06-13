package main

import (
	"fmt"
	"sync"
)

func main() {
	votes := 0
	// mtx := sync.Mutex{}
	wg := &sync.WaitGroup{}

	for range 10 {
		wg.Go(func() {
			for range 10000 {
				// mtx.Lock()
				votes++
				// mtx.Unlock()
			}
		})
	}
	wg.Wait()
	fmt.Println(votes)
}
