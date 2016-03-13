package main

import "fmt"

// START OMIT
func max(x int, y int) int { // HL
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(max(10, 5))
}

//END OMIT
