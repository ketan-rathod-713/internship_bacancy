package main

import "fmt"

func intSeq() func() int { // return func of type return int
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	next := intSeq()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	next2 := intSeq()
	fmt.Println(next2())
	fmt.Println(next2())

	fmt.Println(next())

}
