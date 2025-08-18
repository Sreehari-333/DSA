package main

import "fmt"

func singleNonDuplicate(nums []int) int {

	low := 0
	high := len(nums) - 1

	for low < high {
		mid := low + (high-low)/2

		if nums[mid] != nums[mid+1] || nums[mid] != nums[mid-1] {
			return nums[mid]
		} else if mid%2 == 0 {
			if nums[mid] == nums[mid-1] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] == nums[mid+1] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

	}

	return nums[low]
}

func main() {

	arr := []int{1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	fmt.Println(singleNonDuplicate(arr))
}
