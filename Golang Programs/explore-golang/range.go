package main

import (
	"fmt"
)

// range iterates over elements in a variety of data structures.

func main() {
	slice := []int{1, 2, 3}
	sum := 0
	// index,value for each entry in array, slice
	for _, num := range slice {
		sum += num
	}

	fmt.Println(sum)

	// for map
	// it will also work only for keys too
	mp := map[string]int{"foo": 1, "bar": 2}
	for key, value := range mp {
		fmt.Println(key, value)
	}

}

// DOUBT
// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself. See Strings and Runes for more details
