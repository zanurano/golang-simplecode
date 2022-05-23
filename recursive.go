package main

import (
	"fmt"
)

func reverse(number int, outputNumber int) int {
	// fmt.Println(number)
	if number != 0 {
		outputNumber = outputNumber * 10
		outputNumber = outputNumber + number%10
		number = number / 10
	} else {
		return outputNumber
	}

	return reverse(number, outputNumber)
}

func main() {
	var number int
	fmt.Print("Enter number to reverse ")
	fmt.Scanln(&number)
	fmt.Printf("Reverse of number(%d) is %d\n", number, reverse(number, 0))
}
