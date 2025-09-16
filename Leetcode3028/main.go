package main

import "fmt"

func returnToBoundaryCount(nums []int) int {

	steps := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		steps += nums[i]

		if steps == 0 {
			count++
		}
	}
	return count
}

func main() {

	nums := []int{3, 2, -3, -4}

	fmt.Println(returnToBoundaryCount(nums))
}
