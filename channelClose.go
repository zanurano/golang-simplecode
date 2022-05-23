package main

import (
	"fmt"
	"time"
)

type T int

func isClose(ch <-chan T) bool {
	select {
	case <-ch:
		fmt.Println("close now")
		return true
	default:
	}
	return false
}

func main() {
	c := make(chan T)
	fmt.Println("is close:", isClose(c))
	close(c)
	fmt.Println("is close:", isClose(c))

	quit := make(chan bool)
	n := 0
	go func() {
		fmt.Println("Channel 1")
		for {
			select {
			case <-quit:
				fmt.Println("is done 1")
				return
			default:
				fmt.Println("run", n)
				n++
			}
		}
	}()

	time.Sleep(3000 * time.Millisecond)
	quit <- true
}
