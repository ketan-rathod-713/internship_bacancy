package main

import "fmt"

type Num int
type Str string

func printNum(n Num) {
	// function implementation
	fmt.Println(n)
}

func PrintStr(s Str) {
	// function implementation
	fmt.Println(s)
}

func PrintAny(s interface{}) {
	fmt.Println(s)
}

func main() {
	var num Num = 10
	var str Str = "hello"
	printNum(num)
	PrintStr(str)
	PrintAny(num)
	PrintAny(str)
}
