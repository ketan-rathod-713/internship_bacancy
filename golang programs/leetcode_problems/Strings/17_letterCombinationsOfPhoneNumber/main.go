package main

import (
	"fmt"
)

func main() {
	ans := letterCombinations("23")
	fmt.Println(ans)
}

var mp map[string][]string = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	str := []string{""}

	for _, v := range digits {
		str = mergeArray(str, mp[string(v)])
	}

	return str
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
