package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string
	var c complex128
	fmt.Println(unsafe.Sizeof(str)) // prints 8
	fmt.Println(unsafe.Sizeof(c))   // prints 16

	var a [3]uint32
	fmt.Println(unsafe.Sizeof(a)) // prints 12

	// Structs provide a more flexible way of defining composite types, whose width is the sum of the width of the constituent types, plus padding

	type S struct {
		a uint16
		b uint32
	}
	var s S
	fmt.Println(unsafe.Sizeof(s)) // prints 8, not 6

	var emp struct{}
	fmt.Println(unsafe.Sizeof(emp)) // prints 0

	// you can declare array of structs but it consumes no storage at all
	var x [1000000000]struct{}
	fmt.Println(unsafe.Sizeof(x)) // prints 0

	// Slices of struct{}s consume only the space for their slice header. As demonstrated above, their backing array consumes no space.

	var slice = make([]struct{}, 1000000000)
	fmt.Println(unsafe.Sizeof(slice)) // prints 12 in the playground

	// Interestingly, the address of two struct{} values may be the same.

	var d, b struct{}
	fmt.Println(&b == &d) // true

	// in my case it is false


	// as both the struct doesn't contain any fields hence comparison me dono equla hi aayege.
	e := struct{}{} // not the zero value, a real new struct{} instance
	f := struct{}{}
	fmt.Println(e == f) // true
}
