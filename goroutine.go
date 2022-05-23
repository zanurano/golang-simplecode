package main

import (
	"fmt"
	"math/rand"
)

var cStop = make(chan bool)

func print(till int, message string, ch chan bool) {
	n := 0
	for i := 0; i < till; i++ {
		n += rand.Intn(100)
		fmt.Println((i + 1), message, n)
		if n > 450 {
			fmt.Println("ketemu di", message, n)
			ch <- true

			break
		}
		select {
		case <-ch:
			fmt.Println("is done 1")
			return
			break
		default:
		}

	}

}

func main() {
	// runtime.GOMAXPROCS(2)
	// cStop := make(chan bool)
	go print(10, "aku", cStop)
	go print(10, "kamu", cStop)
	go print(10, "dia", cStop)
	go print(10, "mereka", cStop)
	// print(10, "apa kabar")

	var input string
	fmt.Scanln(&input)
	// fmt.Println(<-cStop)
}
