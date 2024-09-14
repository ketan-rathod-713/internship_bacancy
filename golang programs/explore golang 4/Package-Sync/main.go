package main

import (
	"fmt"
	"sync"
)

func main() {
	once := sync.OnceFunc(func() {
		fmt.Println("I am called")
	})

	once()
	once()
	once()

	
}
