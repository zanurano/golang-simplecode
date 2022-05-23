package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/eaciit/toolkit"
)

func main() {
	jml := 6
	juara := 0
	// ch := make(chan int)

	mx := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(jml)
	for i := 0; i < jml; i++ {
		go func(wt *sync.WaitGroup, x int) {
			defer wt.Done()
			fmt.Printf("%d mulai\n", x)

			jarak := 0
			start := time.Now()
			for {
				mx.Lock()
				if juara >= 3 {
					fmt.Printf("%d berhenti di jarak %d\n", x, jarak)
					mx.Unlock()
					return
				}
				// mx.Unlock()

				jarak += toolkit.RandInt(30)
				mx.Unlock()
				if jarak > 60 {
					mx.Lock()
					// ch <- x
					juara += 1
					durasi := time.Since(start)
					fmt.Printf("%d finish %d durasi %v\n", x, jarak, durasi)
					mx.Unlock()
					return
				}
				time.Sleep(60 * time.Millisecond)
			}
			// time.Sleep(600 * time.Millisecond)
		}(wg, i)
		// time.Sleep(600 * time.Millisecond)
	}

	wg.Wait()
	// fmt.Println("winner", <-ch)
}
