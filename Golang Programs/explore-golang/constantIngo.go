package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var a = 23452345345234523454334
	const a = 23452345345
	// var b int = a
	fmt.Println(reflect.TypeOf(a))
}
