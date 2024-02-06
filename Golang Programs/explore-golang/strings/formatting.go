package main

import "fmt"

func main() {
	n := 103452345253.2

	s := fmt.Sprintf("%5.5f", n)
	fmt.Println(s)
}
