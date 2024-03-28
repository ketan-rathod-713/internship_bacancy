package main

import (
	"fmt"
	"time"
)

// create timer using one channel

func timer(timeout time.Duration) chan int {
	var c chan int = make(chan int)
	// after some time add somethig to this channel
	go func() { // it will call one goroutine which will generate data in channel every second
		for {
			time.Sleep(timeout)
			c <- 1
		}
	}()

	return c
}

func main() {

	// Something that i want to execute each second
	go func() {
		c := timer(time.Second * 1)

		// and then wait for channels each second output to come
		for {
			<-c
			fmt.Println("Timer Of 1 Second")
		}
	}()

	// Something that i want to execute each 2 second
	go func() {
		c := timer(time.Second * 2)

		// and then wait for channels each second output to come
		for {
			<-c
			fmt.Println("Timer Of 2 Second")
		}
	}()

	time.Sleep(10 * time.Second)
}
