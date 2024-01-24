package main

import "fmt"

func main() {

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// This is a good syntax. // sometimes value comes in web request and we just assign it and check.
	if num := 3; num < 10 { // initialization, comparison
		 fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	var err string 
	if err != "" {
		fmt.Println("wow")
	}
}