package main

import (
	"fmt"
	"sync"
)

// You can use WaitGroups to wait for multiple goroutines to finish. A WaitGroup blocks the execution of a function until its internal counter becomes 0.
func main() {
	var wg sync.WaitGroup
	// wait for 2 goroutines at end of main function
	wg.Add(2) // Number of gorutines to wait.

	go goodbye(&wg)
	go HelloWorld(&wg)

	wg.Wait()
}

func goodbye(wg *sync.WaitGroup) {
	fmt.Println("Good bye")
	defer wg.Done()
}

func HelloWorld(wg *sync.WaitGroup) {
	fmt.Println("Hello world")
	defer wg.Done()
}
