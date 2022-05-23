package main

import "fmt"

func main() {
	var word = "Aabcd"
	for _, v := range word {
		// fmt.Println(i, v)
		fmt.Printf("%d %q\n", v, v)
	}

	for _, i := range []int{1, 2, 23, 26} {
		fmt.Printf("%d %q\n", i, toChar(i))
	}
	fmt.Printf("%q %q", xxx(65), xxx(97))
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}
