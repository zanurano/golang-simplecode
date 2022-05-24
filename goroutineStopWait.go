package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	chStop := make(chan struct{})

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, n int, ch <-chan struct{}) {
			defer wg.Done()
			fmt.Printf("Start goroutine %d\n", n)

			select {
			case <-ch:
				fmt.Printf("Stopped goroutine %d\n", n)
			}
		}(wg, i, chStop)
	}

	time.Sleep(3 * time.Second)
	close(chStop)
	fmt.Println("Stopping goroutines 3 second")
	wg.Wait()
	fmt.Printf("all goroutines stopped")
}
