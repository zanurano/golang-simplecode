package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ {
		for j := 0; j < 10; j++ {
			n := 10*i + j
			if n > 107 {
				break
			}
			fmt.Printf("\033[%dm %3d\033[m", n, n)
		}
		fmt.Printf("\n")
	}
}
