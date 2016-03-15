package main

import (
	"fmt"
	"math"
)

// SHAPE OMIT
type Shape interface { // HL
	Area() float64
}

// SHAPEEND OMIT

// SQUARE OMIT
type Square struct{ side float64 } // HL

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

// ENDSQUARE OMIT

// CIRCLE OMIT
type Circle struct{ radius float64 } // HL

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// ENDCIRCLE OMIT

// START OMIT
type Rect struct{ width, height int } // HL

func (r *Rect) Area() int {
	return r.width * r.height
}

func main() {
	var c = &Circle{5}
	var s = &Square{7}
	var r = &Rect{7, 5}

	// boring... call the Area method directly
	fmt.Println(c.Area())
	fmt.Println(s.Area())
	fmt.Println(r.Area())

	// call the Area method via the Shape interface
	// shapes := []Shape{c, s, r}
	// for _, shape := range shapes {
	//	fmt.Println(shape.Area())
	//}
}

// END OMIT
