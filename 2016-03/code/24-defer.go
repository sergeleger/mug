package main

import "os"

// START OMIT
func readFile(filename string) {
	fh, err := os.Open(filename)
	if err != nil {
		return
	}

	defer fh.Close() // HL
}

// END OMIT
