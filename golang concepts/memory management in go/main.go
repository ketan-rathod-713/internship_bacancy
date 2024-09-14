package main

import "fmt"

func main() {
	fmt.Println("creating new duck")
	NewDuck()
}

type Duck struct{}

func NewDuck() Duck {
	return Duck{}
}
