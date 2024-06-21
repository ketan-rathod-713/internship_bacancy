package main

import "fmt"

func main() {
	nums := []int{2, 1, 0, 0, 1, 2}
	sortColors(nums)
}

func sortColors(nums []int) {
	ptr0 := 0
	ptr2 := len(nums) - 1
	fmt.Println(ptr0, ptr2)

	for ptr1 < ptr2 {
		fmt.Println(ptr0, ptr2)
		if nums[ptr0] == 0 {
			ptr0++
			ptr1++
		} else if nums[ptr0] == 2 {
			nums[ptr0], nums[ptr2] = nums[ptr2], nums[ptr0]
            ptr2--
		} else if nums[ptr2] == 2 {
			ptr2--
		} else if nums[ptr2] == 0 {
			nums[ptr0], nums[ptr2] = nums[ptr2], nums[ptr0]
			ptr0++
			ptr1++
		} else { // ptr0 is 1 and ptr2 is 1
			
		}
	}

	fmt.Println(nums)
}
