package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOOS)
	// fmt.Println(runtime.MemProfile())

	// garbage collection is automatic in go
	// if i set somethig to null then it will be collected by garbage collector after some garbage created.
}
