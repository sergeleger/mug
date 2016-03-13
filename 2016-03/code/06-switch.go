package main

import (
	"fmt"
	"time"
)

func today() string {
	return time.Now().Format("Monday")
}

func main() {
	dayOfWeek := today()

	switch dayOfWeek {
	case "Saturday", "Sunday": // HL
		fmt.Println("It's the weekend!")

	default:
		fmt.Println("Get up and go to work.")
	}
}
