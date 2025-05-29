package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {

	left := 0
	right := 1

	sort.Ints(nums)

	for right < len(nums) {
		if nums[left] == nums[right] {
			return true
		}
		left = right
		right++
	}
	return false

}

func main() {

	nums := []int{1, 2, 3, 4, 1}

	fmt.Println(containsDuplicate(nums))
}
