package main

import (
	"fmt"
)

func main() {
	ans := mergeArray([]string{"p", "q"}, []string{"a", "b", "c"})
	fmt.Println(ans)
}

func letterCombinations(digits string) []string {
	ans := []string{""}

	return ans
}

func mergeArray(a, b []string) []string { // min [""] one empty string should be there in an array
	ans := make([]string, 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			concated := fmt.Sprint(a[i], b[j])
			ans = append(ans, concated)
		}
	}

	return ans
}
