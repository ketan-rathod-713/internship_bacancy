package main

import (
	"fmt"
	"reflect"
)

func main() {
	sm := sum(1, 2)
	fmt.Println(sm)

	// TODO variadic function input type
	// multiple return values
	a, b := doubleReturn()
	fmt.Println(a, b)

	// Variadic functions
	// can be called with any number of trailing arguments
	nums := []int{1, 2, 3}
	// this syntax when to pass it as individual arguments not a single slice
	variadicFuncSum(nums...)
	variadicFuncSum(1, 2, 3, 4)

	// it is not allowed to write function definition
	// func sum5(a int, b int){
	// 	fmt.Println("Another method")
	// }

	// Anonymous function, Immediate invoking function
	func() {
		fmt.Println("Wow i am running directly")
	}()

	c, d := namedReturn(12, 10)
	fmt.Println("Returned values are ", c, d)
}

func variadicFuncSum(nums ...int) {
	total := 0
	for _, a := range nums {
		total += a
	}
	fmt.Println(nums, reflect.TypeOf(nums), total)
}

func doubleReturn() (int, int) {
	return 3, 5
}

func sum(a int, b int) int {
	return a + b
}

// auto int to all
func sum2(a, b, c int) int {
	return a + b + c
}

// named return value

func namedReturn(a, b int) (x, y int) {
	fmt.Println("X and Y are initialised already with values ", x, y)
	x = a + b
	y = a * b
	return
}
