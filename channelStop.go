package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//cStopAll := make(chan bool)
	for i := 0; i < 4; i++ {
		fmt.Println("car:", i)
		go func() {
			start := time.Now()
			jarak := 0
			for {
				rnd := rand.Intn(10)
				jarak += rnd
				fmt.Printf("car %d kec %d, tempuh %d\n", i, rnd, jarak)
				if jarak > 50 {
					duration := time.Since(start)
					fmt.Printf("Finish, winner is Car %d, jarak %d, duration %v\n", i, jarak, duration)
					break
				}
				//time.Sleep(20 * time.Millisecond)
			}
		}()
		time.Sleep(300 * time.Millisecond)
	}
}
