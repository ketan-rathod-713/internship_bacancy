package main

import (
	"fmt"
	"time"
)

func main() {
	// don't forget to write make here.
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	time.Sleep(2 * time.Second)
}

// chan pe arrow lagata he, naki variable c pe
func pinger(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}
