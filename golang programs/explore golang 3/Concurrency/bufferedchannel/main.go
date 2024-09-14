package main

import (
	"fmt"
	"sync"
	"time"
)

// use buffered channel here
func main() {
	var c chan string = make(chan string, 5)
	var wg sync.WaitGroup
	wg.Add(2)

	go ping(c, &wg)
	go printer(c, &wg)

	wg.Wait()
}

func printer(c chan string, wg *sync.WaitGroup) {
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout called")
			close(c)
			defer wg.Done()

			// after closing channel check condition
			data, ok := <-c
			if ok {
				fmt.Println("Normal Flow", data)
			} else {
				fmt.Println("Channel is closed")
			}

			return // why break doesn't work here

		}
	}

}

func ping(c chan string, wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		c <- "ping"

	}

	defer wg.Done()
}
