package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	n := 6
	mu := new(sync.Mutex) // lock buat read/write variable "selesai" supaya nggak data race
	selesai := 0          // flag

	wg := new(sync.WaitGroup)
	wg.Add(n) // jumlah goroutine yang harus ditunggu

	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done() // kasih tau waitGroup bahwa goroutine ini sudah selesai
			fmt.Printf("%d mulai berlari\n", i)

			jarak := 0
			for {
				// cek apakah sudah ada yg selesai
				// fmt.Printf("car %d tempuh %d\n", i, jarak)
				mu.Lock()
				if selesai >= 3 {
					fmt.Printf("%d berhenti di jarak %d\n", i, jarak)
					mu.Unlock()
					return
				}
				mu.Unlock()

				jarak += rand.Intn(10)
				if jarak > 50 {
					mu.Lock()
					selesai += 1
					fmt.Printf("%d selesai selelah melewati %d\n", i, jarak)
					mu.Unlock()
					return
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(wg, i)
	}

	wg.Wait() // tunggu semua goroutine selesai sebelum program exit.
}
