package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	right := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[right-1] {
			nums[right] = nums[i]
			right++
		}
	}
	return right
}

func main() {

	arr := []int{1, 1, 2}
	fmt.Println(removeDuplicates(arr))
}
