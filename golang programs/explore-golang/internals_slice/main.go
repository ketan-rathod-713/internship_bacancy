package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	reference := slice[:]

	reference[1] = 20

	fmt.Println(slice)

	// what if i append something to slice
	slice = append(slice, 101)
	reference[2] = 40 // it will not apply as now slice is not referencing that storage location

	fmt.Println(slice)
	fmt.Println(reference)

	var str string = "hello world"
	fmt.Println(string(str[0]))
	// str[0] = 111
}

// Slice internals
// 3 things : len, capacity and pointer to an array.
/*
	type sliceHeader struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}

*/