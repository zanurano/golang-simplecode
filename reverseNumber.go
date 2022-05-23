package main

import "fmt"

var reverse int = 0

func numberReverse2(revnum int) int {
	var remainder int

	if revnum <= 0 {
		return 0

	}
	remainder = revnum % 10
	reverse = reverse*10 + remainder
	numberReverse2(revnum / 10)
	return reverse
}

func numberReverse(revnum int) int {
	var remainder int
	reverse := 0

	for ; revnum > 0; revnum = revnum / 10 {
		remainder = revnum % 10
		reverse = reverse*10 + remainder
	}
	return reverse
}

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		// fmt.Println("remainder:", remainder)
		new_int *= 10
		// fmt.Println("new_int:", new_int)
		new_int += remainder
		// fmt.Println("new_int2:", new_int)
		n /= 10
		// fmt.Println("n:", n)
	}
	return new_int
}

func reverseNumber(num int) int {

	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}

func main() {

	var revnum int

	fmt.Print("Enter the Number to Reverse = ")
	fmt.Scanln(&revnum)

	// reverse := numberReverse(revnum)

	// fmt.Println("The Reverse of the Given Number = ", reverse)

	// reverse2 := numberReverse2(revnum)

	// fmt.Println("The Reverse of the Given Number = ", reverse2)
	fmt.Println("Result:", reverse_int(revnum))

	fmt.Println(reverseNumber(revnum))
}
