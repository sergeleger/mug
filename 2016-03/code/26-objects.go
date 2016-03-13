package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (point Point) Distance(q Point) float64 {
	return math.Hypot(point.X-q.X, point.Y-q.Y)
}

// START OMIT
func (point *Point) ScaleBy(n float64) { // HL
	point.X *= n
	point.Y *= n
}

func main() {
	var p = Point{3, 5}

	fmt.Println("Before:", p)
	p.ScaleBy(3) // HL
	fmt.Println("After:", p)

	dist := p.Distance(Point{1, 0})
	fmt.Println(dist)
}

// END OMIT
