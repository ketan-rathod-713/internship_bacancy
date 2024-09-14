package main

import (
	"fmt"
	"strings"
	"stringsexample/example"
)

func main() {
	a := " amazing great wow"
	// var lastWord string = ""

	// len returns the number of bytes in string
	for i := 0; i < len(a); i++ {
		if strings.Compare(string(a[i]), " ") == 0 {
			fmt.Println("wwo")
			break
		}
	}

	fmt.Println("Example Running")

	// comparison in string :- aab < aaa will return false. Check letter one by one.
	example.Comparison()

	// formatting in string :- using fmt.Sprintf to format float.
	// 20.2%f meaning :- first part is width, second part is precision.
	// For eg. 1.2 will be formatted as "       1.20" as precision is 2 and other 18 spaces are left for matching minimum width.
	// How to pad 0's at the start instead of space : by using %020.2f, 0 before width indicates that we want to pad with 0's instead of space.
	// For more information : https://stackoverflow.com/questions/25637440/how-to-pad-a-number-with-zeros-when-printing

	example.Formatting()
}
