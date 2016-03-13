package main

import "fmt"

func main() {
	capitals := make(map[string]string, 13) // initial space for 13 provinces
	capitals["NB"] = "Fredericton"
	capitals["NS"] = "Halifax"
	capitals["QC"] = "Quebec"
	capitals["ON"] = "Toronto"

	for province, city := range capitals {
		fmt.Printf("%s's capital is %s\n", province, city)
	}

	// Remove an entry from the map
	delete(capitals, "NB")

	fmt.Println(len(capitals))
	fmt.Println(capitals)
}
