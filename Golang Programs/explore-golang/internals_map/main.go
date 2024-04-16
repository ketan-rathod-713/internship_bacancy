package main

import "fmt"

func main() {
	mp := make(map[string]string, 1)

	fmt.Println(len(mp))

	mp["name"] = "ketaan"
	mp["age"] = "22"
	mp["others"] = "nothing"

	fmt.Println(mp, len(mp))
}

// map is always pass by reference
