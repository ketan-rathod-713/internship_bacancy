// fmt package defines a stringer interface // with method String() // for printing any objects stirng representation

// https://pkg.go.dev/fmt#Stringer

package main

import "fmt"

type I interface { // just like stringer interface
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	var i I = a
	val, ok := i.(Person)
	fmt.Println(val, ok)
}

// fmt package look for this interface to print values of specific type
// for ex. for array => bich me space dalna he and start and end me brances ye sab already defined he ha ha.
// do it look for first interface and then to this ?? how 
// TODO: