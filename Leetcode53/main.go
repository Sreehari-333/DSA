package main

import "fmt"

func maxSubArray(nums []int) int {

	currentSum, maxSum := 0, -10000

	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]

		if maxSum < currentSum {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}

func main() {
	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
}
