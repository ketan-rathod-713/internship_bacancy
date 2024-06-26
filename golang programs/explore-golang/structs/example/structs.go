package example

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

func Structs() {
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

	directlyUsingStruct()
}

func directlyUsingStruct() {

	// Here just like bool we can also place a struct here. which would be like anonymous
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
	}

	fmt.Println(s)
}
