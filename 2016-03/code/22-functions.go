package main

import "fmt"

// START1 OMIT
func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(n, m int) int { return n * m }

// START3 OMIT
func main() {
	var f func(int) int

	// assign the square function to f
	f = square
	fmt.Println(f(5))

	// assign the negative function to f
	f = negative
	fmt.Println(f(5))

	// assign the product function to f
	// f = product
	// fmt.Println(f(5, 2))
}

//END1 OMIT
func newProductFunc(n int) func(int) int {
	return func(m int) int {
		return product(n, m)
	}
}

// END3 OMIT
