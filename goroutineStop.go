package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	ch := make(chan string)
	go func() {
		fmt.Println("Mulai..")
		for {
			channel, ok := <-ch
			fmt.Println(!ok)
			if !ok {
				fmt.Println("Shut Down")
				defer wg.Done()
				return
			}
			fmt.Println(channel)
		}
	}()
	time.Sleep(5 * time.Second)
	close(ch)
	wg.Wait()
}
