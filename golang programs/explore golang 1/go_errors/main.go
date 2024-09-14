package main

import (
	"errors"
	"fmt"
	"go_errors/example"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42 ha ha")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
func (e *argError) Error() string {
	return fmt.Sprintf("%v %v", e.arg, e.prob) // why here used Sprintf // because it returns the string instead of writting it on console.
}

func f2(arg int) (int, error) {
	if arg == 42 { // DOUBT - why it requires pointers here.
		return -1, &argError{arg: arg, prob: "can't work with it"}
	} else {
		return arg + 3, nil
	}

}

func main() {
	for _, i := range []int{1, 2, 3, 7, 42, 23} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 succeeded", r)
		}
	}

	fmt.Println("Running Examples")

	example.Example()
	example.Example2()
}
