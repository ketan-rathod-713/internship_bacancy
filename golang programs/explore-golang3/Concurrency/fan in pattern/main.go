package main

import (
	"fmt"
	"time"
)

func producer1() chan int {
	c := make(chan int)
	go func() {
		for {
			c <- 1
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return c
}

func producer2() chan int {
	c := make(chan int)
	go func() {
		for {
			c <- 1
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return c
}

// returns output to multiplexed channel
func fanIn(c1 chan int, c2 chan int) chan int {
	c := make(chan int)

	go func() {
		for {
			select {
			case c <- <-c1:
				fmt.Println("Produce from P1")
			case c <- <-c2:
				fmt.Println("Produce from P2")
			}
		}
	}()

	return c
}

// consume from channel c
func consumer(c chan int) {
	for {
		<-c
		fmt.Println("Consumed")
	}
}

func main() {
	c := fanIn(producer1(), producer2())

	go consumer(c)

	time.Sleep(1 * time.Second)
}

// below output can also occur. It may be look like first it consumed then produced but it is not. it is because we have inserted data in channel before printing it.
/*
Produce from P1
Consumed
Consumed
Produce from P2
Produce from P2
Consumed
Produce from P2
Consumed
Produce from P2
Consumed
*/
