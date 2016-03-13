package main

import "fmt"

// START OMIT
func sum(numbers ...int) (total int) {
	fmt.Printf("%T len=%d\n", numbers, len(numbers))
	for _, v := range numbers {
		total += v
	}
	return
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(2, 4, 6, 8, 10))
}

// END OMIT
