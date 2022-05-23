package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)
	for i := 1; i < 10; i++ {
		<-ticker
		// fmt.Printf("\033[0;0H")
		// fmt.Printf("\x0cOn %d/10", i)
		fmt.Printf("%*s", 20+(i*2), "===>")
		fmt.Printf("\033[0;0H")
		// fmt.Printf("\033[%d;%dH", 5, 0)
	}
	fmt.Println("\nAll is said and done.")
}
