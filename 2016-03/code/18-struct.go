package main

import "fmt"

// START OMIT
type Person struct {
	Name   string
	Street string
	City   string
}

func main() {
	var bob Person
	bob.Name = "Bob"
	bob.Street = "123 Street"
	bob.City = "Somewhere, OH"

	// Using struct literals
	var charles = Person{
		Name:   "Chuck",
		Street: "456 Street",
		City:   "Somecity, ME",
	}

	fmt.Printf("%+v\n", bob)
	fmt.Printf("%+v\n", charles)
}

//END OMIT
