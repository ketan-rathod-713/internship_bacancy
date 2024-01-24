package main

import (
	"fmt"
)

type rect struct {
	width, height int // it will be not exportable
	PublicAttribute int
}

func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
// This methods will work on sructs of type rect only

// Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println("methods in structs")

	r := rect{height: 2, width: 10}

	fmt.Println("area ", r.area())
	fmt.Println("perimeter ", r.perim())

	rp := &r // automatically handles idk what does it mean
	fmt.Println("area ", rp.area())
	fmt.Println("perimeter ", rp.perim())

	hitesh()
}

//  we dont have classes // we have structs so if it contains the function -> methods
func hitesh(){
	fmt.Println("Hitesh methods")
}
