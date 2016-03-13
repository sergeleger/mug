package main

import "fmt"

type Person struct {
	Name   string
	Street string
	City   string
}

// START OMIT
type Employee struct {
	Person // HL
	Id     string
}

// END OMIT

var fmtString = "%-10s\t%-10s\n"

// MAIN OMIT
func main() {
	var bob Employee
	bob.Id = "123" // HL
	bob.Name = "Bob"
	bob.Street = "123 Street"
	bob.City = "Somewhere"

	// Using struct literals
	var charles = Employee{
		Person: Person{Name: "Chuck", Street: "456 Street", City: "Somecity"},
		Id:     "456", // HL
	}

	fmt.Printf(fmtString, "Bob", "Charles")
	fmt.Printf(fmtString, "------------", "------------")
	fmt.Printf(fmtString, bob.Id, charles.Id)         // HL
	fmt.Printf(fmtString, bob.Name, charles.Name)     // HL
	fmt.Printf(fmtString, bob.Street, charles.Street) // HL
	fmt.Printf(fmtString, bob.City, charles.City)     // HL

	// fmt.Printf("%+v\n", bob)
}

// MAINEND OMIT
