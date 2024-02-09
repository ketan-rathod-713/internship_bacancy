package main

import (
	"fmt"
	"math"
)

//An interface type is defined as a set of method signatures.

// A value of interface type can hold any value that implements those methods.

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	// Defining Simple Struct Circle and using its method
	var shape = Circle{radius: 1.0}
	fmt.Println(shape.Area())

	// Above works fine but why address required for below interface Shape
	var s Shape
	s = &Circle{radius: 1} // TODO: Why Need of address here, Else Producing Error
	fmt.Println(s.Area())
}
