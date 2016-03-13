package main

import "fmt"

func main() {
	var capitals = map[string]string{ // HL
		"NB": "Fredericton",
		"NS": "Halifax",
		"QC": "Quebec",
	}

	fmt.Println(capitals["NB"]) // HL
	fmt.Println(capitals["NS"]) // HL

	// Check if a key exists
	city, ok := capitals["ON"] // HL
	if ok {
		fmt.Println("Ontario's capital is", city)
	}
}
