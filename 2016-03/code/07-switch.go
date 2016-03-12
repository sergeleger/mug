package main

import "fmt"

func today() (string, int) {
	return "Tuesday", 2
}

func main() {
	dayOfWeek, weekNumber := today()

	switch {
	case dayOfWeek == "Saturday" || dayOfWeek == "Sunday":
		fmt.Println("It's the weekend!")

	case dayOfWeek == "Tuesday" && weekNumber%2 == 0:
		fmt.Println("Payday!")

	default:
		fmt.Println("Get up and go to work.")
	}
}
