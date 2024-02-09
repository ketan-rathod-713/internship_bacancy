package main

import "fmt"

func main() {
	str := "hello"
	// str2 := []byte(str)
	// fmt.Println(str2)
	// var x uint8
	// x = 104

	var y int32
	y = 104

	// if str[0] == str2[0] {
	// 	fmt.Println("correct")
	// }

	// if str[0] == x {
	// 	fmt.Println("correct")
	// }

	// if str2[0] == x {
	// 	fmt.Println("correct")
	// }

	// https://pkg.go.dev/builtin#byte

	// Iterate over bytes of the string
	for i := 0; i < len(str); i++ {
		fmt.Printf("Byte at index %d: %T %c\n", i, str[i], str[i])
	}

	// https://pkg.go.dev/builtin#rune

	// Convert string to slice of runes and iterate over runes
	runes := []rune(str)
	fmt.Println(runes)
	for i, r := range runes {
		fmt.Printf("Rune at index %d: %T  %d: %c\n", i, r, r, r)
	}

	if runes[0] == y {
		fmt.Println("correct")
	}
}
