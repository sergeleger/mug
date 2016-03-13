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

func (point *Point) ScaleBy(n float64) {
	point.X *= n
	point.Y *= n
}

// START OMIT
type ColoredPoint struct {
	Point // HL
	Color string
}

func main() {
	cpRed := ColoredPoint{Point{1, 2}, "red"}
	cpBlue := ColoredPoint{Point{3, 6}, "blue"}

	cpRed.ScaleBy(3) // ColoredPoint "inherits" the Point methods // HL

	fmt.Println(cpRed.Distance(Point{1, 0}))
	fmt.Println(cpBlue.Distance(Point{1, 0}))
}

// END OMIT
