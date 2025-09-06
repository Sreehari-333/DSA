package main

import "fmt"

func maxProduct(nums []int) int {

	max := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i]-1)*(nums[j]-1) > max {
				max = (nums[i] - 1) * (nums[j] - 1)
			}
		}
	}
	return max
}

func main() {

	nums := []int{3, 4, 5, 2}
	fmt.Println(maxProduct(nums))
}
