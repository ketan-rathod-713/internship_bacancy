package main

import "fmt"

func main() {
	s := "abba"
	ans := lengthOfLongestSubstring(s)
	fmt.Println(ans)
}

func lengthOfLongestSubstring(s string) int {

	ln := len(s)
	if ln == 0 || ln == 1 {
		return ln
	}

	ans := 1 // min 1 to hoga hi

	mp := make(map[rune]int)
	start := -1

	for i, val := range s {
		if v, ok := mp[val]; ok && v > start{
			start = mp[val]// no need to delete map as bad me update hona hi he
		} 

		mp[val] = i // update current val
		fmt.Printf("i %v and start %v\n", i, start)
		ans = max(ans, i - start)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
