package main

import "fmt"

func main() {
	// you want to display 000108 (length = 6)
	// pad zero to total length of 6
	// d
	fmt.Printf("Leading or padded with zero :  |%06d|\n", 108)

	// pad spaces in front(prefix)
	fmt.Printf("Leading or padded with space : |%6d|\n", 108)

}
