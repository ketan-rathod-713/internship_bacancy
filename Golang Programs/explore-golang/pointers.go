package main

import "fmt"

func zeroVal(a int) {
	a = 0
}

func zeroPtr(a *int) {
	*a = 0
}

func main() {
	fmt.Println("weolcome to pointers")

	var a int = 10
	zeroVal(a)
	fmt.Println(a)

	zeroPtr(&a)
	fmt.Println(a)

}
