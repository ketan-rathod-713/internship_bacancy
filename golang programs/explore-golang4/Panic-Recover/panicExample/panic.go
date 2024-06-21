package panicexample

import "fmt"

func PanicRecoverExample() {
	defer func() {
		fmt.Println("exit normally.")
	}()
	fmt.Println("hi!")
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()
	panic("bye!")

	// How to make this code reachable even though i am using panic and recover
	fmt.Println("unreachable")
}