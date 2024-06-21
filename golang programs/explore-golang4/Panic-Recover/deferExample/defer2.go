// eval-moment.go
package deferExample

import "fmt"

func DeferInsideFunction() {
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i+x)
		}
		x = 10
	}()

	// 0 1 2 // prints 2 1 0

	// defer functions parameter values are evaluated at the time of function call.
	fmt.Println()
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			// This will work as closure function in go 1.22
			defer func() {
				fmt.Println("b:", i+x)
			}()
		}
		x = 10
	}()
}
