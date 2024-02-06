package main

import (
	"fmt"
	"reflect"
)

var (
	d = 30
	e = 50
)

func main() {
	// var a = 23452345345234523454334
	const a = 23452345345
	// var b int = a
	fmt.Println(reflect.TypeOf(a))

	var ( // to delcare multiple variables
		m = 30
		n = 50
	)

	fmt.Println(m, n)
}
