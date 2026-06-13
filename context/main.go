package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo is finished")
			return
		default:
			fmt.Println("Foo is working")
		}

		time.Sleep(400 * time.Millisecond)
	}
}

func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo is finished")
			return
		default:
			fmt.Println("Boo is working")
		}

		time.Sleep(400 * time.Millisecond)
	}
}

func goo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goo is finished")
			return
		default:
			fmt.Println("Goo is working")
		}

		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	middleCtx, middleCancel := context.WithCancel(parentCtx)
	childCtx, childCancel := context.WithCancel(middleCtx)

	go foo(parentCtx)
	go boo(middleCtx)
	go goo(childCtx)

	time.Sleep(time.Second)
	parentCancel()
	time.Sleep(time.Second)
	middleCancel()
	time.Sleep(time.Second)
	childCancel()
	time.Sleep(time.Second)
}
