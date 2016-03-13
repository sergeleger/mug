package main

import (
	"fmt"
	"math/rand"
	"time"
)

func coinFlip() string {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(100)
	if n < 40 {
		return "heads"
	} else if n < 80 {
		return "tails"
	} else {
		return ""
	}
}

func main() {
	var heads, tails int

	switch coinFlip() { // HL
	case "heads": // HL
		heads++

	case "tails": // HL
		tails++

	default: // HL
		fmt.Println("landed on edge!")
	}

	fmt.Println(heads, tails)
}
