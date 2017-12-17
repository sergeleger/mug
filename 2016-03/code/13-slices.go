package main

import "fmt"

func main() {
	ids := make([]int, 5) // HL
	ids[0] = 2
	ids[1] = 15
	fmt.Println("len=", len(ids), "cap=", cap(ids))
	fmt.Println(ids)
}
