package main

import (
	"fmt"
	"log"
)

// START OMIT
func sum(numbers ...int) (total int) {
	defer log.Printf("Deferred? %T len=%d\n", numbers, len(numbers)) // HL

	for _, v := range numbers {
		defer log.Println(v, total) // HL
		total += v
	}

	log.Println("leaving sum()")
	return
}

func main() {
	fmt.Println(sum(2, 4, 6, 8, 10))
}

// END OMIT
