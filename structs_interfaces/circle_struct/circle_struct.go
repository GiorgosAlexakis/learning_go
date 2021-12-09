package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

/*
alternatively:
type Circle struct {
	x, y, r float64
}
*/

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	// var c Circle ,creates a local variable c of type Circle that
	// each value is set to their corresponding zero value
	// c:= new(Circle), same as above but returns a pointer to the struct (*Circle)
	// using new like this is uncommon
	c := Circle{0, 0, 5}
	// c := &Circle{0,0,5} if you want a pointer to the struct
	fmt.Println(c.x, c.y, c.r)
	fmt.Println("Circle Area with Radius 5: ", circleArea(c))
}
