package main

import "fmt"

func main() {
	var nums = []int{2, 1, 3, 4}
	var k = 1

	fmt.Println("Final Ans Is", minOperations(nums, k))
}

func minOperations(nums []int, k int) int {
	var steps int = 0
	var output int = 0
	for _, num := range nums {
		output = num ^ output
	}

	fmt.Println(output)

	for {
		fmt.Println("Output:", output, "K:", k)
		if output % 2 == k % 2 {
			// then no need to add to steps
		} else {
			steps++
		}

		// if both are 0 then break
		if output == k && k == 0 {
			break;
		}

		output = output >> 1
		k = k >> 1

	}

	return steps
}
