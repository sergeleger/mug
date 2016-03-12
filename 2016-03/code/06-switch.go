package main

import "fmt"

func today() string {
	return "Tuesday"
}

func main() {
	dayOfWeek := today()

	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")

	default:
		fmt.Println("Get up and go to work.")
	}
}
