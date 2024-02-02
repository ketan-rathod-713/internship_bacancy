package main

import (
	"fmt"
	"math"
)

func main() {
	var input string = "(1+(2*3)+((8)/4))+1"
	var ans int = maxDepth(input)
	fmt.Println("The max depth is ", ans)
}

func maxDepth(s string) int {
	var open int = 0

	var maxDepth int = 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			open++
		} else if string(s[i]) == ")" {
			maxDepth = int(math.Max(float64(open), float64(maxDepth)))
			open--
		}
	}

	return maxDepth
}
