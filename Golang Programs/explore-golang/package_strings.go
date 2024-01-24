package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Lets explore strings package")

	str := "this, is, the, string"

	fmt.Println(strings.Count(str, "is"))
	arr := strings.Split(str, ",")
	fmt.Println(len(arr))
}
