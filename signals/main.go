package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT)
	workersCount, err := strconv.Atoi(os.Getenv("WORKERS_COUNT"))
	if err != nil {
		panic(err)
	}

	for i := range workersCount {
		wg.Go(func() {
			ticker := time.NewTicker(2 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Bye from Worker", i)
					return
				case <-ticker.C:
					fmt.Println("Message from Worker", i)
				}
			}
		})
	}

	wg.Wait()
}
