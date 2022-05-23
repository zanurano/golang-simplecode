package main

import "testing"

type TestNumberData struct {
	Number  int
	Reverse int
}

func TestReverseNumber(t *testing.T) {
	testReverses := []TestNumberData{
		{12, 21},
		{212, 212},
		{1234, 4321},
		{12456, 65421},
		{1000, 1},
		{1100, 11},
	}

	for _, val := range testReverses {
		res := ReverseNumber(val.Number)
		if res != val.Reverse {

		}
	}
}
