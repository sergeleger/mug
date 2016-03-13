package main

import (
	"errors"
	"fmt"
)

// START OMIT
func addProduct(x, y int) (int, int) {
	return x + y, x * y
}

func safeAddProduct(x, y int) (int, int, error) {
	if x < 0 || y < 0 {
		return 0, 0, errors.New("can't work with negative numbers")
	}
	return x + y, x * y, nil
}

func main() {
	a, b := addProduct(2, 5)
	fmt.Println(a, b)

	x, y, err := safeAddProduct(-2, 5)
	if err != nil {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println(x, y)
	}
}

//END OMIT
