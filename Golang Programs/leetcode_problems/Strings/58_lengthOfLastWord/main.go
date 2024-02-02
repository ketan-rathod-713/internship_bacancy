package main

import (
	"strings"
)

func main() {
	var input string = "  fly me   to   the moon  "
	var ans int = lengthOfLastWord(input)
	println(ans)
}

// DOUBT
// How this function works ? TrimSpace
// https://cs.opensource.google/go/go/+/refs/tags/go1.21.6:src/strings/strings.go;l=1007

func lengthOfLastWord(input string) int {
	input = strings.TrimSpace(input)

	var ln int = 0
	for _, val := range input {
		if string(val) == " " {
			ln = 0
			continue
		}
		// else increase the length
		ln++
	}

	return ln
}
