package main

import "fmt"
import "os"
import "io"

// START OMIT
type myReader uint8

func (r myReader) Read(buf []byte) (int, error) { // io.Reader "implementation" // HL
	for i := range buf {
		buf[i] = byte(r)
	}
	return len(buf), nil
}

func main() {
	var reader myReader = 65

	// Read 10 bytes from the reader...
	var buf = make([]byte, 10)
	reader.Read(buf)

	fmt.Println(buf)

	io.Copy(os.Stdout, reader)
}

// END OMIT
