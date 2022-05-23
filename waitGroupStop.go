package main

import (
	"fmt"
	"sync"
)

func worker(messages <-chan int, shutdown <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case message := <-messages:
			//Do something useful with message here
			fmt.Println("ok", message)
		case <-shutdown:
			//We're done!
			fmt.Println("done")
			return
		}
	}
}

func main() {
	messages := make(chan int)
	shutdown := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go worker(messages, shutdown, wg)
	close(shutdown) //Signal to shutdown
	wg.Wait()       //Wait for the goroutine to shutdown
}
