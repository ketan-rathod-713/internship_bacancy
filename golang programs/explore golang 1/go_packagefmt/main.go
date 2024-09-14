package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Printf("iit prints values %v %v and so on \n", a, b)
	fmt.Printf("iit prints Types of values %T and %T \n", a, b)
	fmt.Printf("iit prints base 2 value %b and %b \n", a, b)
	fmt.Printf("iit prints base 10 value %d and %d \n", a, b)

	c := 12.4567
	fmt.Printf("%.2f \n\n", c)
}
