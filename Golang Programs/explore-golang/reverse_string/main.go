package main

import "fmt"

func main() {
	ans := reversePrefix("goodtosee", 'd')

	fmt.Println(ans)
}

func reversePrefix(word string, ch byte) string {
	// tranverse from left to right for ch byte
	// if it is there then mark it as point and break from loop

	byteArr := []byte(word)
	var point int = -1

	for i, char := range word {
		if byte(char) == ch {
			point = i
			break
		}
	}

	if point == -1 {
		return word
	} else {
		// reverse given string slice of string
		mid := point / 2
		for i := 0; i <= mid; i++ {
			temp := byteArr[i]
			byteArr[i] = byteArr[point - i]
			byteArr[point - i] = temp
		}

		return string(byteArr)
	}
}
