package main

import "fmt"

// import "C"

func main() {
	fmt.Printf("%s\n", "yess")
	fmt.Printf("%*s\n", 30, "000")
	fmt.Printf("I'm a %[1]s, a great %[2]s\n", "programmer", "zanur")

	fmt.Print("\033c")
	fmt.Print("\033c\n")
	fmt.Print("\x1bc")
	fmt.Print("\x1bc\n")
	fmt.Println("\033c")
	fmt.Println("\x1bc")

	fmt.Printf("%s\n", "yess")
	fmt.Printf("%*s\n", 30, "000")
	fmt.Printf("I'm a %[1]s, a great %[2]s\n", "programmer", "zanur")

	fmt.Println(7 % 8)
}
