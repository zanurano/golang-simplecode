package main

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func main() {
	var number int
	fmt.Print("Enter number to check is palindrom: ")
	fmt.Scanln(&number)
	// number := 21321

	// arrNumber := toArray(number)
	// fmt.Println(arrNumber)

	// fmt.Println(mergeToInt(arrNumber))
	if IsPalindrome(number) {
		fmt.Println(number, "Is Palindrom", Green(IsPalindrome(number)))
	} else {
		fmt.Println(number, "Is Palindrom", Red(IsPalindrome(number)))
	}

	// number = 1000
	// if IsPalindrome(number) {
	// 	fmt.Println(number, "Is Palindrom", Green(IsPalindrome(number)))
	// } else {
	// 	fmt.Println(number, "Is Palindrom", Red(IsPalindrome(number)))
	// }
}

func ReverseNumber(number int) int {
	res := 0

	for number != 0 {
		temp := number % 10
		// fmt.Println(temp)
		res = (res + temp) * 10
		// fmt.Println("res:", res)

		number /= 10
	}

	// fmt.Println(res / 10)

	return res / 10
}

func toArray(number int) []int {
	var array []int

	for number != 0 {
		array = append(array, number%10)
		number /= 10
	}

	return array
}

func mergeToInt(s []int) int {
	hasil := 0
	nominal := 1
	for i := len(s) - 1; i >= 0; i-- {
		hasil += s[i] * nominal
		nominal *= 10
	}
	return hasil
}

func IsPalindrome(num int) bool {
	/*
		1. Tidak boleh diconvert ke string atau rune
		2. hanya menggunakan standard library
	*/
	// arrNumber := toArray(num)
	// numberReverse := mergeToInt(arrNumber)
	numberReverse := ReverseNumber(num)

	if num == numberReverse {
		return true
	}
	return false
}
