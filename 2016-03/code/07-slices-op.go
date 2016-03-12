package main

import "fmt"

func main() {
	var daysOfWeek = [7]string{
		"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday",
		"Saturday",
	}

	weekdays := daysOfWeek[1:6]
	fmt.Println(weekdays)

	// But the weekdays slice is backed by the same array
	weekdays[0] = "Raining Mondays"
	fmt.Println(daysOfWeek)
}
