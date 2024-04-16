package main

import (
	"fmt"
	"sliceexample/example"
)

func main() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := arr[3:5]
	fmt.Println(slice)

	fmt.Println("Address of slice ", &slice)
	fmt.Println("capacity and length of slice is ", cap(slice), len(slice))
	sliceReturned := passingSliceToFunction(&slice)
	fmt.Println(arr)
	fmt.Println("slice returned ", sliceReturned)
	fmt.Println("capacity and length of slice is ", cap(slice), len(slice))
	fmt.Println("Address of slice ", &slice)

	fmt.Println("")
	fmt.Println("Length and capacity")
	lengthAndCapOfSliceSame()

	fmt.Println("")
	fmt.Println("Example of copy function")
	exampleOfCopyFunc()

	fmt.Println("Example running")

	example.SliceExample()
	example.SliceExample2()
}

func lengthAndCapOfSliceSame() {
	slice := make([]int, 10)
	fmt.Println(len(slice), cap(slice)) // 10, 10
	slice2 := slice[1:]
	fmt.Println(len(slice2), cap(slice2)) // 9 9

	// Now what was my doubbt yes
	slice2 = slice2[0:5]
	fmt.Println(len(slice2), cap(slice2)) // it will decrease the length to 5 and capacity remainded 9

}

func exampleOfCopyFunc() {
	slice := make([]int, 10)
	fmt.Println(slice)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3

	// now copy from specified index
	copy(slice[3:], slice[0:])
	// copy first 3 elements to next 3 elements in slice ha ha

	fmt.Println(slice)
}

// Slice Header stores
// type sliceHeader struct {
//     Length        int
//     Capacity      int
//     ZerothElement *byte
// }

func passingSliceToFunction(slice *[]int) []int { // 0-5 slice pointer
	// (*slice)[0] = 50

	// it will change the slice header, slice header is copied here which is pointing to the array elements. hence this will not reflect in the original array ha ha
	*slice = append(*slice, 30, 40, 50, 60)
	*slice = (*slice)[0:7] // rather then append we can use this but it will give error when index out of capacity // hence use append instead.
	// what happens after append
	(*slice)[6] = 100
	return *slice
}

// capacity increase //

// Even though the slice header is passed by value, the header includes a pointer to elements of an array, so both the original slice header and the copy of the header passed to the function describe the same array. Therefore, when the function returns, the modified elements can be seen through the original slice variable.
// https://go.dev/blog/slices

// Here we see that the contents of a slice argument can be modified by a function, but its header cannot. The length stored in the slice variable is not modified by the call to the function, since the function is passed a copy of the slice header, not the original. Thus if we want to write a function that modifies the header, we must return it as a result parameter, just as we have done here. The slice variable is unchanged but the returned value has the new length, which is then stored in newSlice,
