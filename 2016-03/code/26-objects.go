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
func (point *Point) ScaleBy(n float64) {
	point.X *= n
	point.Y *= n
}

func main() {
	var p = Point{3, 5}

	fmt.Println(p)
	p.ScaleBy(3)
	fmt.Println(p)

	fmt.Println(p.Distance(Point{1, 0}))
}

// END OMIT
