package main

import (
	"fmt"
	"strings"
)

func main() {
	a := " amazing great wow"
	// var lastWord string = ""

	for i := 0; i < len(a); i++ {
		if strings.Compare(string(a[i]), " ") == 0 {
			fmt.Println("wwo")
			break
		}
	}
}
