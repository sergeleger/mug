package main

import (
	"errors"
	"fmt"
)

// START OMIT
func addAndProduct(x, y int) (int, int) { // HL
	return x + y, x * y
}

func safeAddAndProduct(x, y int) (int, int, error) { // HL
	if x < 0 || y < 0 {
		return 0, 0, errors.New("can't work with negative numbers")
	}
	return x + y, x * y, nil
}

func main() {
	a, b := addAndProduct(2, 5)
	fmt.Println(a, b)

	// x, y, err := safeAddAndProduct(-2, 5)
	// if err != nil {
	// 	fmt.Println("error occured:", err)
	// } else {
	// 	fmt.Println(x, y)
	// }
}

//END OMIT
