package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 4, 5, 6}

	variadic(nums...)
}

func variadic(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}
