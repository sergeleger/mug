package main

import "fmt"

var dwarves = [7]string{
	"Doc",
	"Grumpy",
	"Happy",
	"Sleepy",
	"Bashful",
	"Sneezy",
	"Dopey",
}

func main() {
	// C style loop
	for i := 0; i < 7; i++ {
		fmt.Println(dwarves[i])
	}

	fmt.Println("-----")

	// Using the range operator
	for index, name := range dwarves {
		fmt.Println(index, name)
	}
}
