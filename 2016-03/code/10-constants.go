package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("Sunday = ", Sunday)
	fmt.Println("Saturday = ", Saturday)
}
