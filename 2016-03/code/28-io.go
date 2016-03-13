package main

// START OMIT

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// Interface embedding // HL
type ReadCloser interface {
	Reader
	Closer
}

// END OMIT
