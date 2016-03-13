package main

import "fmt"

type Person struct {
	Name   string
	Street string
	City   string
}

// START OMIT
type Employee struct {
	Person
	Id string
}

func main() {
	var bob Employee
	bob.Id = "123"
	bob.Name = "Bob"
	bob.Street = "123 Street"
	bob.City = "Somewhere, OH"

	// Using struct literals
	var charles = Employee{
		Person: Person{Name: "Chuck", Street: "456 Street", City: "Somecity, ME"},
		Id:     "456",
	}

	fmt.Printf("%+v\n", bob)
	fmt.Printf("%+v\n", charles)
}

//END OMIT
