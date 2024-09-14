package practice

import (
	"fmt"
	"time"
)

func TimeoutWithSelect() {
	c := something()

	for {
		select {
		case <-c:
			fmt.Println("i am running")
		case <-time.After(time.Microsecond * 4):
			fmt.Println("You are too late ha ha !!")
			return
		}
	}
}

func something() chan string {
	c := make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			time.After(time.Second * 5)
			c <- "good bro"
			time.After(time.Second * 5)
		}()
	}

	return c
}

func routine(input1, timeout chan string) <-chan string {
	c := make(chan string)

	select {
	case v1 := <-input1:
		fmt.Println("case 1 done", v1)
	case <-timeout:
		fmt.Println("case 2 done")
	}

	return c
}

func DaisyChain() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}

func f(left, right chan int) {
	left <- 1 + <-right // when left and right both are ready then this execute
}
