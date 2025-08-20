package main

import "fmt"

func transformArray(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i] = 0
		} else {
			nums[i] = 1
		}
	}
	// return nums

	for j := 0; j < len(nums); j++ {
		for k := j + 1; k < len(nums); k++ {
			if nums[j] > nums[k] {
				temp := nums[j]
				nums[j] = nums[k]
				nums[k] = temp
			}
		}
	}
	return nums
}

func main() {
	nums := []int{4, 3, 2, 1}
	fmt.Println(transformArray(nums))
}
