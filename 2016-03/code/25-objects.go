package main

import (
	"fmt"
	"math"
)

// START OMIT
type Point struct {
	X, Y float64
}

func (point Point) Distance(q Point) float64 {
	return math.Hypot(point.X-q.X, point.Y-q.Y)
}

func main() {
	var p = Point{3, 5}

	fmt.Println(p.Distance(Point{1, 0}))
}

// END OMIT
