package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "0P"// check if palindrome, should return true
	fmt.Println(isPalindrome(input))
}

func isAlphaNumeric(char string) bool{
	if((char >= "a" && char <= "z") || (char >= "A" && char <= "Z") || (char >= "0" && char <= "9")){
		return true;
	}

	return false;
}

// don't encounter any space for calculation
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	fmt.Println(s)

	// logic for with space string
	var i int = 0
	var j int = len(s) - 1

	for i < j{
		// also remove non alpha numeric
		if(!isAlphaNumeric(string(s[i]))){
			i++
			continue
		}
		if(!isAlphaNumeric(string(s[j]))){
			j--
			continue
		}

		fmt.Println("Values are ", string(s[i]), " and ", string(s[j]))
		if(s[i] == s[j]){
			i++
            j--
			continue
		} else {
			return false;
		}
	}

	
	return true;
}
