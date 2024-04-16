package main

import (
	"fmt"
	"time"
)

func Merge(out chan int, a chan int, b chan int) {
	for {
		select {
		case v := <-a:
			out <- v
		case v := <-b:
			out <- v
		}
	}
}

func main() {
	var out chan int = make(chan int, 0)
	var a chan int = make(chan int, 0)
	var b chan int = make(chan int, 0)
	go Merge(out, nil, nil)
	go func() {
		a <- 1
		b <- 3
		b <- 3
		b <- 3
		b <- 3
		b <- 3
		b <- 3
		b <- 3
		b <- 3
		b <- 3
		b <- 3
		b <- 3
	}()
	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()

	// for v := range out {
	// 	fmt.Println(v)
	// }
	time.Sleep(3 * time.Second)

	// nil pointers and interfaces
	var t *tree
	var c Calculate = t
	fmt.Println(c == nil) // false because it is interface with decrete type and withotu value
	fmt.Println(t == nil)
}

type tree struct {
}

func (t tree) Sum() int {
	return 0
}

type Calculate interface {
	Sum() int
}
