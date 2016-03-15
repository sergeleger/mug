package main

import (
	"fmt"
	"math/rand"
	"time"
)

func someFunction() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func main() {
	percent := someFunction()
	if percent < 50 { // HL
		fmt.Println("You're just starting.", percent)
	} else { // HL
		fmt.Println("Almost done.", percent)
	}
}
