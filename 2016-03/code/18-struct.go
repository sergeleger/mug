package main

import (
	"encoding/json"
	"fmt"
)

var _ = json.Marshal

// START OMIT
type Person struct {
	Name   string
	Street string
	City   string
}

// END OMIT

var fmtString = "%-10s\t%-10s\n"

// MAIN OMIT
func main() {
	var bob Person // HL
	bob.Name = "Bob"
	bob.Street = "123 Street"
	bob.City = "Somewhere"

	// Using struct literals
	var charles = Person{ // HL
		Name:   "Chuck",
		Street: "456 Street",
		City:   "Somecity",
	}

	fmt.Printf(fmtString, "Bob", "Charles")
	fmt.Printf(fmtString, "------------", "------------")
	fmt.Printf(fmtString, bob.Name, charles.Name)     // HL
	fmt.Printf(fmtString, bob.Street, charles.Street) // HL
	fmt.Printf(fmtString, bob.City, charles.City)     // HL

	// fmt.Printf("%+v\n", bob)
}

// MAINEND OMIT
