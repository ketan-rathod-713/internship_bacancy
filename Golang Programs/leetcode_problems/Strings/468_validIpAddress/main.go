package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ans := validIPAddress("192.0.0.1")
	fmt.Println(ans)
}

func validIPAddress(queryIP string) string {
	// first try for ipv4 only
	arrayOfStrings := strings.Split(queryIP, ".")
	arrayOfStrings2 := strings.Split(queryIP, ":")
	
	if(len(arrayOfStrings) == 4){
		if(checkIpV4(arrayOfStrings)){
			return "IPv4"
		}
	} else if(len(arrayOfStrings2) == 8){
		// fmt.Println(arrayOfStrings2)
		if(checkIpv6(arrayOfStrings2)){
			return "IPv6"
		}
	} 

	return "Neither"
}

func checkIpv6(arrayOfStrings []string) bool {
	for i:=0; i<len(arrayOfStrings); i++ {
		if(checkValidIpv6Number(arrayOfStrings[i])){
			// fmt.Println(arrayOfStrings[i])
			continue
		} else {
			return false
		}
	}
	return true
}

func checkValidIpv6Number(str string) bool {
	if(len(str) <= 0 || len(str) > 4){
		// fmt.Println("String length is not appropriate for ", str, len(str))
		return false
	}
	// ceck if it contains alphanumerics else return false
	for i := 0; i < len(str); i++ {
		if((string(str[i]) >= "a" && string(str[i]) <= "f") || (string(str[i]) >= "A" && string(str[i]) <= "F" || (string(str[i]) >= "0" && string(str[i]) <= "9"))){
			// do nothing
		} else {
			return false
		}
	}
	return true
}
func checkIpV4(arrayOfStrings []string) bool{
	for i := 0; i < len(arrayOfStrings); i++ {
		if(checkValidNumber(arrayOfStrings[i])){
			continue
		} else {
			return false
		}
	}

	return true
}

func checkValidNumber(str string) bool {
	ans, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false
	} else {
		if(ans == 0 && len(str) == 1){
			return true
		} else
		if (ans >= 0 && string(str[0]) == "0") || ans > 255 {
			return false
		}

		return true
	}
}
