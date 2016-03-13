package main

import "fmt"

// START OMIT
func describe(v interface{}) { // HL
	fmt.Printf("(%T, %v)\n", v, v)
}

func main() {
	var f32, f64 = float32(21), float64(42)
	var i, s = 10, "A string"

	describe(f32)
	describe(f64)
	describe(i)
	describe(&i)
	describe(s)
}

// END OMIT
