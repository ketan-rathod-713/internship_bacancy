package main

import "fmt"

func main() {
	var s string = "paper"
	var t string = "title"

	ans := isIsomorphic(s, t)
	fmt.Println(ans)
}

// each time i make error in this ha ha
// error test case : "bdac" and "baba"
// means dono side se check karna padega

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mp1 := make(map[string]string) // left to right
	mp2 := make(map[string]string) // left to right

	for i := 0; i < len(s); i++ {
		// map 1 uniqueness
		if mp1[string(s[i])] == "" { // if no values entered // then can add value here // can we add directly yes ig
			mp1[string(s[i])] = string(t[i])
		} else { // check if both values are equal
			if mp1[string(s[i])] == string(t[i]) {
				// do nothing
			} else {
				return false
			}
		}

		// map 2 uniqueness
		if mp2[string(t[i])] == "" { // if no values entered // then can add value here // can we add directly yes ig
			mp2[string(t[i])] = string(s[i])
		} else { // check if both values are equal
			if mp2[string(t[i])] == string(s[i]) {
				// do nothing
			} else {
				return false
			}
		}

		// fmt.Println(mp1, mp2)
	}
	return true
}
