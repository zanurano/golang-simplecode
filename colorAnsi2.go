package main

import "fmt"

type (
	Red    string
	Green  string
	Yellow string
	Blue   string
)

func (r Red) Format(f fmt.State, c rune) {
	f.Write([]byte("\033[31m" + string(r) + "\033[0m"))
}
func (g Green) Format(f fmt.State, c rune) {
	f.Write([]byte("\033[32m" + string(g) + "\033[0m"))
}
func (y Yellow) Format(f fmt.State, c rune) {
	f.Write([]byte("\033[33m" + string(y) + "\033[0m"))
}
func (b Blue) Format(f fmt.State, c rune) {
	f.Write([]byte("\033[34m" + string(b) + "\033[0m"))
}
func main() {
	var (
		b Blue   = "Hel"
		y Yellow = "lo "
		g Green  = "Wor"
		r Red    = "ld!"
	)
	fmt.Println(b, y, g, r)

	fmt.Println(Blue("Connection closed."))
}
