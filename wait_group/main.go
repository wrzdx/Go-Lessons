package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func waterGarden(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	n := (5 + rand.Intn(6)) * 100
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Println("Garden is watered by", number)
}

func main() {
	wg := &sync.WaitGroup{}
	for i := range 1 + rand.Intn(5) {
		wg.Add(1)
		go waterGarden(wg, i)
	}

	wg.Wait()
	fmt.Println("Watering is done")
}
