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
	} else if n < 45 {
		return "tails"
	} else {
		return ""
	}
}

func main() {
	var heads, tails int
	switch coinFlip() {
	case "heads":
		heads++

	case "tails":
		tails++

	default:
		fmt.Println("landed on edge!")
	}

	fmt.Println(heads, tails)
}
