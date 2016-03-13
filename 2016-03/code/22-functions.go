package main

import "fmt"

// START1 OMIT
func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(n, m int) int { return n * m }

// START3 OMIT
func main() {
	var fn func(int) int // HL

	// assign the square function to f
	fn = square // HL
	fmt.Println(fn(5))

	// assign the negative function to f
	fn = negative // HL
	fmt.Println(fn(5))

	// assign the product function to f
	// fn = product
	// fmt.Println(fn(5, 2))
}

//END1 OMIT
func newProductFunc(n int) func(int) int { // HLnew
	return func(m int) int { // HLnew
		return product(n, m) // HLnew
	} // HLnew
} // HLnew

// END3 OMIT
