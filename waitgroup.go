package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				println("done")
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	wg.Wait()

	// var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for b := range c {
			fmt.Printf("Hello %s\n", b)
		}
		println("done")
	}()
	c <- "zanur"
	c <- "omar"
	c <- "aisyah"
	close(c)
	wg.Wait()
}
