package main

import (
	"fmt"
	"math"
)

func main() {
	text1 := "abc"
	text2 := "dabc"
	lcs := longestCommonSubsequence(text1, text2, 0, 0)
	fmt.Println(lcs)
}

// TIME LIMIT EXCEEDED // USE DP
func longestCommonSubsequence(text1 string, text2 string, i int, j int) int {
	// fmt.Println("lcs called with", text1, text2,i, j)
	l1 := len(text1)
	l2 := len(text2)

	condMatchedResult := 0
	cond1Result := 0
	cond2Result := 0

	if(i == l1 || j == l2){
		return 0; // yaha se kuch nahi ho payega ha ha
	} else if(text1[i] == text2[j]){ // if cureent char matched then add 1 and go on searching for more
		condMatchedResult = 1 + longestCommonSubsequence(text1, text2, i+1, j+1)
	} else { // not matched then 2 condtions
		cond1Result = longestCommonSubsequence(text1, text2, i+1, j)
		cond2Result = longestCommonSubsequence(text1, text2, i, j+1)
	}

	mx := int(math.Max(float64(condMatchedResult), math.Max(float64(cond1Result), float64(cond2Result))))
	return mx
}
