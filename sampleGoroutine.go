package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	jml := 4
	wg := new(sync.WaitGroup)
	// wg.Add(jml)
	//cStopAll := make(chan bool)
	for i := 0; i < jml; i++ {
		wg.Add(1)
		go func(wt *sync.WaitGroup, n int) {
			defer wt.Done()
			fmt.Printf("mobile %d jalan\n", n)

			start := time.Now()
			jarak := 0
			for {
				rnd := rand.Intn(25)
				jarak += rnd
				// fmt.Printf("car %d kec %d, tempuh %d\n", n, rnd, jarak)
				if jarak > 50 {
					duration := time.Since(start)
					fmt.Printf("Finish, winner is Car %d, jarak %d, duration %v\n", n, jarak, duration)
					break
				}
				time.Sleep(200 * time.Millisecond)
			}
		}(wg, i)
		// time.Sleep(30 * time.Millisecond)
	}

	wg.Wait()
}
