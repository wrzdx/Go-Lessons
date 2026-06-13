package main

import (
	"fmt"
	"sync"
)

func main() {
	mails:=[]string{}
	mtx := sync.Mutex{}
	wg := &sync.WaitGroup{}

	for range 10 {
		wg.Go(func() {
			for range 10000 {
				mtx.Lock()
				mails = append(mails, "mail")
				mtx.Unlock()
			}
		})
	}
	wg.Wait()
	fmt.Println(len(mails))
}
