package main

import "fmt"

func main() {
	var input string = "(()())(())"
	ans := removeOuterParentheses(input)
	fmt.Println("Ans is ", ans)
}

func removeOuterParentheses(s string) string {
	var ans string = ""
	var open int = 0
	var close int = 0;

	for i:=0; i<len(s); i++ {
		if string(s[i]) == "("{
			open++
		}
		if string(s[i]) == ")"{
            close++
        }

		// Remove

		if(open == 1 && close == 0){
			// do nothing
		} else if(open == close){
			// do nothing
			open = 0
			close = 0
		} else { // not a parent parenthesis
			ans += string(s[i])
		}
	}


	return ans
}
