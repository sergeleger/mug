package main

import (
	"fmt"
	"log"
)

// DOUBLESTART OMIT
func double(x interface{}) interface{} {
	if i, ok := x.(int); ok {
		return i * 2
	} else if iPtr, ok := x.(*int); ok {
		return (*iPtr) * 2
	} else if f32, ok := x.(float32); ok {
		return f32 * 2
	} else if f64, ok := x.(float64); ok {
		return f64 * 2
	} else {
		log.Printf("unsupported type: %T\n", x)
	}

	panic("unsupported type")
}

// DOUBLEEND OMIT

// DOUBLESWITCHSTART OMIT
func double2(x interface{}) interface{} {
	switch x := x.(type) {
	case int:
		return x * 2
	case float32:
		return x * 2
	case float64:
		return x * 2
	case *int:
		return *x * 2
	default:
		log.Printf("unsupported type: %T\n", x)
	}

	panic("unsupported type")
}

// DOUBLESWITCHEND OMIT

// MAIN OMIT
func main() {
	fmt.Println(double(float32(5.5)))
	fmt.Println(double(float64(5.5)))

	i := 15
	fmt.Println(double(i))
	fmt.Println(double(&i))
}

// MAINEND OMIT
