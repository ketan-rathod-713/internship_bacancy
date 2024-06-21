package main

import (
	"fmt"
)

func main() {
	input := []string{"dog", "racecar", "car"}
	ans := longestCommonPrefix(input)
	fmt.Println("common prefix is ", ans)
}

// LOGIC
// Check One by one each and at the end of array append to ans
// check if input string size is not exceeded if exceeded then return common_str there only
func longestCommonPrefix(input []string) string {
	var common_str string

	if len(input[0]) == 0 {
		return common_str
	}

	var commonChar string
	for i := 0; i < len(input[0]); i++ { // number of passes, increasing length each time

		if i < len(input[0]) {
			commonChar = string(input[0][i])
		} else {
			return common_str
		}

		//fmt.Println("Common char for pass ", i, " is ", commonChar)

		for j := 0; j < len(input); j++ { // iterate for each input string
			//fmt.Println(j, i, string(input[j][i]))

			if i >= len(input[j]) {
				//fmt.Println("Size overload ", i, " ", len(input[j]))
				return common_str // can't do anything now
			}

			if commonChar != string(input[j][i]) {
				//fmt.Println("Not equal ", commonChar, " ", string(input[j][i]))
				return common_str
			}

		}
		// if everything is fine till end of the pass then append it to common_str
		common_str += string(input[0][i])
	}

	return common_str
}
