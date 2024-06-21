package main

import "fmt"

var (
	a = b
	b = initialiseB()
	c = 3
)

func initialiseB() int {
	return c + 2
}

func init() {
	println("init function called")
}

func main() {

	// This part is giving me error
	// var (
	// 	e = g
	// 	g = h
	// 	h = 2
	// )
	// it will not work here

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
