package main

import "fmt"

func main() {
	for _, c := range []int{1, 3, 65, 97} {
		fmt.Printf("%q\n", toChar(c))
	}

	fmt.Printf("%q", xxx(65))
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func xxx(x rune) rune {
	return x
}
