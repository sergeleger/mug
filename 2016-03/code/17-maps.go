package main

import "fmt"

func main() {
	capitals := make(map[string]string, 13) // initial space for 13 provinces // HL
	capitals["NB"] = "Fredericton"
	capitals["NS"] = "Halifax"
	capitals["QC"] = "Quebec"
	capitals["ON"] = "Toronto"

	for province, city := range capitals { // HL
		fmt.Printf("%s's capital is %s\n", province, city)
	}

	// Remove an entry from the map
	delete(capitals, "NB") // HL

	fmt.Println(len(capitals)) // HL
	fmt.Println(capitals)
}
