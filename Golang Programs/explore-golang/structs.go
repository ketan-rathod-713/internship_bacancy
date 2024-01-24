package main

import (
	"fmt"
)

// person struct type has name and age field
type person struct {
	name string
	age  int
}

// creates new person and returns it
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 43
	return &p // safely return the pointer
}

func main() {
	fmt.Println("structs in golang")
	// no inheritance in golang; no super or parent

	// creating new struct
	fmt.Println(person{name: "ketan", age: 23})

	fmt.Println(newPerson("jhon"))

	s := newPerson("aman")
	fmt.Println(s.name)

	// structs are mutable
	s.name = "something"

	// anonymous structs are also valid
}
