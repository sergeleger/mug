package main

import "fmt"

type Color uint64

func main() {
	var red Color = 1
	fmt.Println(red)

	// Same rules apply, you cannot assign a Color value to an uint64 variable
	// var value uint64
	// value = red
}
