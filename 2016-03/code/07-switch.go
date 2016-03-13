package main

import (
	"fmt"
	"time"
)

func today() string {
	return time.Now().Format("Monday")
}

func week() int {
	_, n := time.Now().ISOWeek()
	return n
}

func main() {
	dayOfWeek := today()
	weekNumber := week()

	switch {
	case dayOfWeek == "Saturday" || dayOfWeek == "Sunday": // HL
		fmt.Println("It's the weekend!")

	case dayOfWeek == "Tuesday" && weekNumber%2 == 0: // HL
		fmt.Println("Payday!")

	default:
		fmt.Println("Get up and go to work.")
	}
}
