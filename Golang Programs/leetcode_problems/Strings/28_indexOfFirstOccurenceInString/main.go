package main

import "fmt"

func main() {
	var ans int = strStr("mississippi", "issipi")
	fmt.Println(ans)
}

// check if needle is present in haystack
// if present then return start position
// else return -1
func strStr(haystack string, needle string) int {

	if len(needle) > len(haystack) {
		return -1
	}

	// now iterate all char of haystack
	for i := 0; i < len(haystack); i++ {
		// iterate for each items of haystack
		currI := i

		for j := 0; j < len(needle); {
			fmt.Println(currI, j, string(haystack[currI]), string(needle[j]))
			// check if haystack out of bound condition
			if currI >= len(haystack) {
				return -1
			}
			if needle[j] == haystack[currI] {
				currI++
				j++ // no need to increase it
			} else {
				break // if not equal then break here
			}
			if j == len(needle) {
				return i
			}
		}
	}

	return -1
}
