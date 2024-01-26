package main

import "fmt"

type Info struct {
	name string
	age  int
}

type Animal struct {
	Info
	kind string
}

func (a Animal) Eat() {
	fmt.Println("Animal", a.Info.name, a.Info.age, " is eating")
}

func main() {
	fmt.Println("Welcome to struct embeddings")

	animal := Animal{
		Info: Info{
			name: "<NAME>",
			age:  3,
		},
		kind: "Dog",
	}

	animal.Eat()
}
