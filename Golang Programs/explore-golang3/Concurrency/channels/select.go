package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select Example")

	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	go func ()  {
		for {
			c1 <- "ping"
		time.Sleep(1 * time.Second)
		}
	}()

	go func(){
		for {
			c2 <- "pong"
			time.Sleep(3 * time.Second)
		}
	}()

	go func(){
		// dono mese jobhi pehle aa jaye usko lelo 
		for {
			select {
			case msg := <-c1:
				fmt.Println(msg)
			case msg := <-c2:
				fmt.Println(msg)
			case <-time.After(1 * time.Second):
				fmt.Println("Timeout called")
			}
		}
	}()

	time.Sleep(20 * time.Second)
}
