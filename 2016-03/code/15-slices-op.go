package main

import (
	"fmt"
	"strings"
)

func main() {
	var daysOfWeek = [7]string{
		"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
	}

	// Slice from an array
	weekdays := daysOfWeek[1:6] // HL

	// Range operator over a slice
	for i, v := range weekdays {
		fmt.Println(i, v)
		weekdays[i] = strings.ToLower(v)
	}

	fmt.Println(weekdays)

	// But the weekdays and daysOfWeek slices are backed by the same array:
	// fmt.Println(daysOfWeek)
}
