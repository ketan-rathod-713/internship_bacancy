package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 10
	}()

	fmt.Println("got", <-c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	cancel()
	Something(ctx)
}

func Something(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Function cancelled")
	case <-time.After(2 * time.Second):
		fmt.Println("After 2 second delay i give output")
	}
}
