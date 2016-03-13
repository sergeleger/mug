package main

import (
	"fmt"
	"time"
)

// START OMIT
func say(str string) {
	fmt.Println(str)
}

func main() {
	go say("Hello, World!")
	say("All Done!")

	time.Sleep(500 * time.Millisecond)
}

// END OMIT
