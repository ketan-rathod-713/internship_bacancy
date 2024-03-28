
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled, stopping long running task")
			return
		default:
			fmt.Println("Working on long running task...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go longRunningTask(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("Canceling long running task...")
	cancel()

	time.Sleep(1 * time.Second) // Wait for the task to finish
}