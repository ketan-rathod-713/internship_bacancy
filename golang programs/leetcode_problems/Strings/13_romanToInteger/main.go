package main

import "fmt"

func main() {
	var input []string = []string{"III", "IV", "XL", "LVIII", "MCMXCIV"}

	for _, str := range input {
		ans := romanToInt(str)
		fmt.Println("Ans for ", str, " is ", ans)
	}
}

// SOLUTION
/*
6 instances where substraction is used.
IX - 9
IV - 4
XL - 40
XC - 90
CD - 400
CM - 900

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000


*/

/*
 check bigger one first and add directly as small ke bad big may be aa sakte he

*/

func romanToInt(s string) int {
	charValue := make(map[string]int)
	charValue["I"] = 1
	charValue["V"] = 5
	charValue["X"] = 10
	charValue["L"] = 50
	charValue["C"] = 100
	charValue["D"] = 500
	charValue["M"] = 1000

	charPriority := make(map[string]int)
	charPriority["I"] = 1
	charPriority["V"] = 2
	charPriority["X"] = 3
	charPriority["L"] = 4
	charPriority["C"] = 5
	charPriority["D"] = 6
	charPriority["M"] = 7

	// Now main logic

	var sum int = 0

	for i := 0; i < len(s); {
		if i+1 < len(s) {
			// check priority
			if charPriority[string(s[i])] >= charPriority[string(s[i+1])] { // directly add it and move on to next
				sum += charValue[string(s[i])]
				i++
			} else { // means there is need of substraction
				sum += (charValue[string(s[i+1])] - charValue[string(s[i])])
				// Two times ++ as two characters are calculated here
				i++
				i++
			}
		} else {
			sum += charValue[string(s[i])]
			i++
		}
	}

	return sum
}
