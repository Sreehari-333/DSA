package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mid = 0
	for start <= end {
		mid = start + (end-start)/2

		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(search(nums, 6))
}
