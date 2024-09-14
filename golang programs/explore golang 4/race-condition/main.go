package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var i int64 = 0

var wg sync.WaitGroup

func incr() {
	defer wg.Done()
	atomic.AddInt64(&i, 1)
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incr()
	}

	wg.Wait()
	fmt.Println("Value of i is", i)
}
