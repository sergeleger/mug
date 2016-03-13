package main

import (
	"fmt"
	"io"
)

// START OMIT
type myReader uint8

func (r myReader) Read(p []byte) (int, error) { // io.Reader "implementation" // HL
	max := 100

	for i := range p {
		p[i] = byte(r)
		if i+1 == max {
			return max, io.EOF
		}
	}
	return len(p), nil
}

func main() {
	var reader myReader = 65

	// Read 10 bytes from the reader...
	var buf = make([]byte, 10)
	reader.Read(buf)
	fmt.Println(buf)

	// io.Copy(os.Stdout, reader)
}

// END OMIT
