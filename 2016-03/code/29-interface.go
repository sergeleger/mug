package main

import (
	"fmt"
	"math"
)

// SHAPE OMIT
type Shape interface {
	Area() float64
}

// SHAPEEND OMIT

// SQUARE OMIT
type Square struct{ side float64 }

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

// ENDSQUARE OMIT

// CIRCLE OMIT
type Circle struct{ radius float64 }

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// ENDCIRCLE OMIT

// START OMIT
type Rect struct{ width, height int }

func (r *Rect) Area() int {
	return r.width * r.height
}

func main() {
	var c = &Circle{5}
	var s = &Square{7}
	var r = &Rect{7, 5}

	fmt.Println(c.Area())
	fmt.Println(s.Area())
	fmt.Println(r.Area())
}

// END OMIT
