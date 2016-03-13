package main

import "fmt"

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

func main() {
	fmt.Println("Sunday = ", Sunday)
	fmt.Println("Saturday = ", Saturday)
}

// Final Version:
// const (
// 	Sunday = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// )
