package main

import "fmt"

func main() {
	values := []string{"bird", "5", "animal"}

	// Pad all values to 10 characters.
	// ... This right-justifies the strings.
	//     Three periods just for decoration.
	for i := range values {
		fmt.Printf("%10v...\n", values[i])
	}

	// Pad all values to 10 characters.
	// ... This left-justifies the strings.
	//     Vertical bars just for decoration.
	for i := range values {
		fmt.Printf("|%-10v|\n", values[i])
	}
}
