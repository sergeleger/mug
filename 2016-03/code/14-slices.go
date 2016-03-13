package main

import "fmt"

func main() {
	var ids = make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		fmt.Println(len(ids), cap(ids))
		ids = append(ids, i)
	}
}
