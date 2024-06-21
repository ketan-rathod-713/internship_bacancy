package main

import "fmt"

func main() {
	fmt.Println("Hellow Arrays")

	fmt.Println("why it is running faster idk may be its ubuntu thats why ha ha")

	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// by default arrays have 0 value
	var arr2 [5]int
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
}
