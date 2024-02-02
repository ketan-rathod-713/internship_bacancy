package main

import "fmt"

type struct KeyValuePair {
	Key string 
	Value string
}

func main() {
	input := "trrrree" // output should be eert

	// Solution // Store in map and then take out items one by one

	mp := make(map[string]int)
	for i := 0; i < len(input); i++ {
		mp[string(input[i])]++
	}

	fmt.Println(mp)
	// now remove items based on their values

	// map is not sorted directly hence convert it to array of structs
	keyValuePairs := []KeyValuePair{}
	for key, value := range mp {
		keyValuePairs = append(keyValuePairs, KeyValuePair{key, fmt.Sprintf("%d", value)})
	}

	// Define a sorting function
	sort.Slice(keyValuePairs, func(i,j  int) bool {
		return keyValuePairs[i].Value < keyValuePairs[j].Value
	})
}
