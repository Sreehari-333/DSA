package main

import "fmt"

func majorityElement(nums []int) int {
	count := 0
	var candidate int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
