package main

import (
	"fmt"
	"io"
	"os"
)

type myReader uint8

func (r myReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = byte(r)
	}
	return len(p), nil
}

// START OMIT
func readAndClose(r io.Reader) {
	// ... read data from r ...

	// Close the file if it implements the io.Closer interface
	if closer, ok := r.(io.Closer); ok { // Type Assertion to another interface // HL
		fmt.Println("Closing the file")
		closer.Close()
	}

	// Is this a myReader?
	if _, ok := r.(myReader); ok { // Type Assertion to the implementation type // HL
		fmt.Println("You've been using a dumb reader.")
	}
}

func main() {
	var reader myReader = 65
	readAndClose(os.Stdin)
	readAndClose(reader)
}

// END OMIT
