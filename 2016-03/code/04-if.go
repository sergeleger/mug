package main

import (
	"fmt"
	"math/rand"
)

func someFunction() int {
	return rand.Intn(100)
}

func main() {
	percent := someFunction()
	if percent < 50 {
		fmt.Println("You're just starting.")
	} else {
		fmt.Println("Almost done.")
	}
}
