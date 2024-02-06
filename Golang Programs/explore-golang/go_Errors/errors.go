package main

// Go programs express error state with error values.

// The error type is a built-in interface similar to fmt.Stringer:

// type error interface {
//     Error() string
// }

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) String() string {
	return fmt.Sprintf("at string %v, %s",
		e.When, e.What)
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error { // error use kar pa rahe he because MyError has already implemented error by implicitly adding Error() method to it
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
