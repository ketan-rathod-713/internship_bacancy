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

	learnPointer()

}

type Vertex struct {
	X int
	Y int
}

func learnPointer() {
	var a int = 10
	var b *int = &a
	*b = 20
	fmt.Println(a)

	v := Vertex{1, 2}
	p := &v
	fmt.Println((*p).X, (*p).Y, " and ", p.X, p.Y, " Both are same ha ha")
}
