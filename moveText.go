package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker := time.Tick(time.Second)
	// for i := 1; i <= 10; i++ {
	// 	<-ticker
	// 	fmt.Printf("\x0cOn %d/10", i)
	// 	fmt.Printf("%*s\n", 20+i, "==>")
	// 	fmt.Printf("\033[0;0H")
	// }
	// fmt.Println("\nAll is said and done.")
	// for i := 10; i >= 0; i-- {
	// 	fmt.Printf("\033[2K\r%d", i)
	// 	time.Sleep(1 * time.Second)
	// }
	// fmt.Println()
	ticker := time.Tick(time.Second)
	for i := 1; i <= 10; i++ {
		<-ticker
		// fmt.Printf("\x0cOn %d/10", i)
		fmt.Printf("%*s\n", 20+i, "==>")
		fmt.Printf("\033[%d;%dH", 3, 1)

	}
	fmt.Println("\nAll is said and done.")
}
