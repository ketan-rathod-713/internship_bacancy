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

// Interfaces are implemented implicitly we don;t need to specify it

// package main

// import "fmt"

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// // This method means type T implements the interface I,
// // but we don't need to explicitly declare that it does so.
// func (t T) M() {
// 	fmt.Println(t.S)
// }

// func main() {
// 	var i I = T{"hello"}
// 	i.M()
// }

// interface values
// (value, type) // stores this value under the hood
