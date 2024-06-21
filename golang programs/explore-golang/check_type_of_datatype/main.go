package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "this is string"

	dataType := reflect.TypeOf(str)
	fmt.Println("data type of ", str, " is ", dataType)

	var a int = 20
	fmt.Println("dt of a is ", reflect.TypeOf(a))

	fmt.Println("by using reflect.TypeOf(str) or using string formatting %T ig")
}
