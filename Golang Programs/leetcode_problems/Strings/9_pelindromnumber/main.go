package main

import "fmt"

func main() {
	var x = 1012
	ans := isPalindrome(x)
	fmt.Println(ans)
}

func isPalindrome(x int) bool {
	str := fmt.Sprint(x)
	i := 0
	j := len(str) - 1

	for i < j {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}
