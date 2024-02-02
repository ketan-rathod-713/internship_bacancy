package main

import "fmt"

func main() {
	var ans int = findMin([]int{3,1,2})
	fmt.Println(ans)
}

	func findMin(nums []int) int {
		n := len(nums)
		low, high := 0, n-1

		for low < high {
			mid := low + (high-low)/2
			if(nums[mid] > nums[high]){ // result must be in right half
				low = mid + 1
			} else if(nums[mid] <= nums[high]){ // result may be this or in left half
				high = mid
			} 
		} 

		return nums[low]
	}

// Time Limit Exceeded Here
// func findMin(nums []int) int {
// 	start := 0
// 	end := len(nums) - 1
// 	ans := 10

// 	startVal := nums[start]
// 	endVal := nums[end]
// 	mid := start + (end-start)/2
// 	for start <= end {

// 		if nums[mid] > startVal && nums[mid] > endVal { // first and last dono se bada then go right
// 			// go right
// 			start = mid + 1
// 		} else if nums[mid] <= endVal {
// 			ans = nums[mid]
// 			end = mid - 1
// 		}

// 		mid = start + (end-start)/2
// 	}

// 	return ans
// }
