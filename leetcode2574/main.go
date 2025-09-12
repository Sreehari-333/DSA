package main

import "fmt"

func leftRightDifference(nums []int) []int {

	// answers := []int{}
	leftSum := []int{}
	rightSum := make([]int, len(nums))
	lSum := 0

	for i := 0; i < len(nums); i++ {
		leftSum = append(leftSum, lSum)
		lSum += nums[i]
	}

	rSum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		rightSum[i] = rSum
		rSum += nums[i]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if leftSum[i] > rightSum[i] {
			result[i] = leftSum[i] - rightSum[i]
		} else {
			result[i] = rightSum[i] - leftSum[i]
		}
	}

	return rightSum
}

func main() {

	nums := []int{10, 4, 8, 3}
	fmt.Println(leftRightDifference(nums))
}
