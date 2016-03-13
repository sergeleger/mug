package main

import (
	"io"
	"os"
)

// START OMIT
type myReader uint8

func (r myReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = byte(r)
	}
	return len(p), nil
}

func main() {
	var reader myReader = 65
	io.Copy(os.Stdout, reader)

	// &io.LimitedReader{reader, 10}
}

// END OMIT
