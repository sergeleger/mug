package main

import "fmt"

// START OMIT
func sum(numbers ...int) (total int) {
	defer fmt.Printf("Deferred? %T len=%d\n", numbers, len(numbers))
	for _, v := range numbers {
		defer fmt.Println(v, total)
		total += v
	}

	fmt.Println("leaving sum()")
	return
}

func main() {
	fmt.Println(sum(2, 4, 6, 8, 10))
}

// END OMIT
