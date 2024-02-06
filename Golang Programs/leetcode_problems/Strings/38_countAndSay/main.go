package main

import "fmt"

func main() {
	ans := countAndSay(4)
	fmt.Println(ans)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	// else calculate using previous one
	prevString := countAndSay(n - 1)
	// fmt.Println(n, prevString)
	return processPrevString(prevString)
}

func processPrevString(s string) string {
	var start string
	count := 0
	processedString := ""

	for i, val := range s {
		if i == 0 {
			start = string(val)
		}

		if start == string(val) {
			count++
		} else {
			processedString += fmt.Sprint(count, start)
			count = 1 // count this atleast and reset start
			start = string(val)
		}
	}

	// process last part
	processedString += fmt.Sprint(count, start)

	return processedString
}
