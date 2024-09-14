package main

import (
	"bufio"
	"bytes"
	"fmt"
	"runtime"
)

func main() {
	var buf [8192]byte
	n := runtime.Stack(buf[:], false)

	sc := bufio.NewScanner(bytes.NewReader(buf[:n]))

	for sc.Scan() {
		var p uintptr
		n, _ := fmt.Sscanf(sc.Text(), "testing")
		if n != 1 {
			continue
		}

		fmt.Println("returning", p)
	}
}
