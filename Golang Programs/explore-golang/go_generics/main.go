package main

import "fmt"

// We can define methods on generic types just like we do on regular types, but we have to keep the type parameters in place. The type is List[T], not List.

func main() {
	fmt.Println("Has a value ", Has([]string{"a", "b", "c"}, "a"))
	fmt.Println("Has a value ", Has([]string{"a", "b", "c"}, "d"))

	fmt.Println("Has a value ", HasGeneric([]string{"a", "b", "c"}, "a"))
	fmt.Println("Has a value ", HasGeneric([]string{"a", "b", "c"}, "d"))

	fmt.Println("Has a value ", HasGeneric([]int{1, 2, 3, 4}, 2))
	fmt.Println("Has a value ", HasGeneric([]int{1, 2, 3, 4}, 10))

	fmt.Printf("My list %v ", NewEmptyList[int]()) // if inputs in the type parameters doesn't specify all then it will cause an error
	// fmt.Printf("My list %v ", NewEmptyList[]()) // It will cause an error

	// Multiple type parameters
	PrintThings(10, 20, 30, 40) // No error here
	// PrintThings(10, 20, "wow", "great") // It will give an error

	PrintThings[int, string, int](10, 20, "a", 3) // this will work as A and B can be of any type while C can only be of type int ha ha

	UsingGenericTypes()

	fmt.Println("Understanding struct map")
	understandStructMap()
}

func PrintThings[A, B any, C ~int](a1, a2 A, b B, c C) {
	fmt.Printf("\n %v %v %v %v \n", a1, a2, b, c)

}

func Has(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

// Internal logic of Has will be same as that of other data types

// before parameters list there will be type parameters list in square brackets.

// we give each type parameter a name and followed by a constraint

func HasGeneric[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

// compiler does type interence based on function calls

func NewEmptyList[T any]() []T {
	return make([]T, 0)
}

// we can write our own generic type
// Set - unique value

// Set[T] means map[T]struct{} // with empty value
// type parameters and then it contains the type definition

// iska matlab ye he ki internally set map me store hoga
type Set[T comparable] map[T]struct{}

// generic function 
func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
    return set;
}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	fmt.Println("inside Has function value and ok are ",value, ok)
    return ok
}

// Simple example

// Pair is a generic type that holds two values of any type.
// In this way we define generic type
type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// let me have generic Type of string

// it can hold the slice of Type Any T
type StringGeneric[T any] []T

func UsingGenericTypes() {
	pair := Pair[int, int]{1, 2} // this is a generic type which we can use anywhere ha ha
	fmt.Println(pair)

	stringArray := StringGeneric[string]{} // no need to put here [] because it is just like {} for array values also.
	fmt.Println("generic array of any type ", stringArray)

	intSet := NewSet(1,2,3)
	fmt.Println("", intSet)
	fmt.Println(intSet.Has(2))
	fmt.Println(intSet.Has(5))
}

// // To define the node of the linked list
// type Node[T1 string, T2 Node] struct {
// 	name T1
// 	next *Node[T1, T2]
// }


func understandStructMap(){
	set := make(map[string] struct {})
	
	set["a"] = struct{}{} // it declares an empty struct// size is not allocated for it.
	fmt.Println(set)
	fmt.Println(set["a"])
	fmt.Println(set["a"])
}

