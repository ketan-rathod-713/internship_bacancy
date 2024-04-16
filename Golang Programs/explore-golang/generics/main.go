package main

import "fmt"

type Num int

// It will accept Type that is alias for int or int
func Add[T ~int](a T, b T) {
	fmt.Println(a + b)
}

type CustomData interface {
	int | float64
}

type User[T CustomData] struct {
	Id   int
	Name string
	Data T
}

func main() {
	var a Num = 10
	var b Num = 10

	Add(a, b)

	// Hare we need to specify what type we are using, else it will give an error ha ha
	u := User[int]{Id: 1, Name: "string", Data: 12}
	fmt.Println(u)
}

// For map key must be comparable
