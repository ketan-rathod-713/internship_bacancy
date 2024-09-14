package main

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

// Ping Pong Pattern
func main() {
	table := make(chan *Ball)

	// start three goroutines ping and pong
	go player("manav", table)
	go player("ketan", table)
	go player("vatsal", table)
	// go player("more", table)

	// add ball to channel

	table <- new(Ball) // game on; toss the ball
	time.Sleep(1 * time.Second)

	// remove ball from channel
	<-table // game over

	// panic("show me the stacks")
}

func player(name string, table chan *Ball) {

	// it will try to grab the ball and increase the balls hits and then sleep and return back to table
	for {
		ball := <-table
		// It is essential to put below logic before adding something to channel
		ball.hits++
		fmt.Println(name, ball.hits)

		table <- ball
	}
}
