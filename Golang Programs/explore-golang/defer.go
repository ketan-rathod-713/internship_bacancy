package main

import "fmt"

func main() {
	// Last in first out pattern for defer.
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")


	defer fmt.Println("Three")

	fmt.Println("World")

	myDefer()
}

func myDefer(){
	for i:=0;i<10;i++ {
		defer fmt.Println(i)
	}
}