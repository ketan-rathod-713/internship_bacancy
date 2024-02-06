package main

import "fmt"

func main() {
	slice1 := []int{1, 3}
	slice2 := []int{2, 4}

	slice3 := [][]int{slice1, slice2}
	fmt.Println(slice3)

	slice1[0] = 30
	slice2[1] = 40
	slice2 = append(slice2, 10)
	slice2[1] = 90

	fmt.Println(slice3)
}
