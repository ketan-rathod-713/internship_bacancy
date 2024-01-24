// An uninitialized slice equals to nil and has length 0.
// to create empty slice -> make

package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	slice := make([]string, 3)
	slice[0] = "string"
	slice[1] = "is"
	slice[2] = "good"

	fmt.Println(slice, len(slice), cap(slice))

	// append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value

	slice = append(slice, "with", "append")
	fmt.Println(slice)

	// copy slices // no error if small
	slice2 := make([]string, 7)
	copy(slice2, slice)
	fmt.Println(slice2, slice2[6], reflect.TypeOf(slice2[6])) // string obviusly

	if slice2[6] == "" {
		fmt.Println("empty string at slice2[6]")
	}

	// check dt of uniitialised values -. string

	// slice // low and high // high is excluded
	fmt.Println(slice[1:3])
	fmt.Println(slice[1:])
	fmt.Println(slice[:2]) // excluding 2 here

	// Declare and Initialise slice in single line
	// DOUBT
	// isn't it array syntax same ?? but ig that was fixed and it is not
	slice3 := []int{1, 2, 3}
	fmt.Println(slice3, reflect.TypeOf(slice3))

	slice4 := []int{1, 2, 4}
	fmt.Println(slice4)

	slice4 = append(slice4, 5, 6)
	fmt.Println(slice4)

	hitesh()
}

func hitesh() {
	var fruitList = []string{"apple", "tomato", "peach", "something"}

	// fruitList = append(fruitList[1:])

	fruitList = append(fruitList, "something ha ha")

	fmt.Println(fruitList)

	fmt.Println(sort.StringSlice(fruitList)) // not working ig
	sort.Strings(fruitList)                  // It will work now
	fmt.Println(fruitList)

	highscore := []int{10, 20, 10, 4, 3}
	sort.Ints(highscore)

	fmt.Println(highscore)

	fmt.Println(sort.IntsAreSorted(highscore)) // it will return true as highscores is sorted.

	// removing element from slice based on index

	courses := []string{"string", "python", "java", "switf", "rust"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After deleting second index ", courses)
}
