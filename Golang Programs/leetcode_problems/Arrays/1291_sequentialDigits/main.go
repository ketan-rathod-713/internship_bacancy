package main

import (
	"fmt"
)

func main() {
	ans := sequentialDigits(100, 300) // 123 234 and till <= 300
	fmt.Println(ans)
}

// length of low // first sequential digit find // 150 then 3 digits then
func sequentialDigits(low int, high int) []int {

	initialLength = len(fmt.Sprint(low))

	// Now try for each iteration

	str := createSeqDigit(3, 1)
	fmt.Println(str)

	ans := []int{}

	return ans
}

func createSeqDigit(ln int, start int) string {
	str := ""
	for i := 0; i < ln; i++ {
		str += fmt.Sprint(start)
		start++
	}

	return str
}
