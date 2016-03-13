package main

import (
	"fmt"
	"log"
)

// START OMIT
func sum(numbers ...int) (total int) { // HL
	log.Printf("%T len=%d\n", numbers, len(numbers))

	for _, v := range numbers {
		total += v
	}
	return
}

func main() {
	fmt.Println("Sum is", sum())
	fmt.Println("Sum is", sum(2, 4, 6, 8, 10))
}

// END OMIT
