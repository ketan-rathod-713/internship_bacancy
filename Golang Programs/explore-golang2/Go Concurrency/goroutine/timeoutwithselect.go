package goroutine

import (
	"fmt"
	"time"
)

/*

time.After function returns a channel that blocks for specified duration, after the interval, the channel delivers the current time, once.

*/

func TimeoutMain() {
	c := boring2("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(400 * time.Millisecond):
			fmt.Println("YOu are too slow")
			return
		}
	}
}

func TimeoutWholeConversionUsingSelectMain() {
	// create timer once, outside the loop to timeout entire conversation.
	// In previous program we had time out for each message.
	c := boring2("Timeout")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Your timeout is done ha ha nikklyo  sab cpu ke
			bahar")
			return
		}
	}
}

/* Quit Channel
25:50


*/


/*  System Software 
Go made for making system softwares at google.
Design Fake framework 	TODO:

Example of google search engine

*/