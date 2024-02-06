package main

import "fmt"

// TODO what to do
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // There is no type hence it will produce a run time error // not compile time ha ha
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
