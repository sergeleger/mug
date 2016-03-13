package main

import "fmt"

func main() {
	var capitals = map[string]string{
		"NB": "Fredericton",
		"NS": "Halifax",
		"QC": "Quebec",
	}
	fmt.Println(capitals["NB"])
	fmt.Println(capitals["NS"])

	// Check if a key exists
	city, ok := capitals["ON"]
	if ok {
		fmt.Println("Ontario's capital is", city)
	}
}
