package main

import "fmt"

func main() {
	var daysOfWeek = [7]string{
		"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday",
		"Saturday",
	}

	fmt.Println("Number of entries: ", len(daysOfWeek))
	fmt.Println("First day:", daysOfWeek[0])
	fmt.Printf("%T", daysOfWeek)
}
