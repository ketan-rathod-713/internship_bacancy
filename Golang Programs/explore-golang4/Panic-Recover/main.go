package main

import (
	"internals/deferExample"
	panicexample "internals/panicExample"
)

func main() {
	deferExample.Defer()

	deferExample.DeferInsideFunction()

	panicexample.PanicRecoverExample()

}

// Example of gorutine block time error
// package main

// var c chan int

// func waiter() {
// 	for {

// 	}
// 	<-c
// }

// func main() {
// 	c = make(chan int)

// 	go waiter()

// 	c <- 10

// 	// time.Sleep(1 * time.Second)

// 	// fmt.Println("Now its time to close")
// }
