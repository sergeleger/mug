package main

import "fmt"

const boillingF = 212.0

func main() {
	var f = boillingF
	c := (f - 32) * 5/9
	fmt.Printf("boilling point = %f°F or %g°C\n", f, c)
}
