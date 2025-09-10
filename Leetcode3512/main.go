package main

import "fmt"

func minOperations(nums []int, k int) int {

	numOfOperations := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for sum%k != 0 {
		sum--
		numOfOperations++
	}
	return numOfOperations
}

func main() {

	nums := []int{3, 9, 7}
	k := 5
	fmt.Println(minOperations(nums, k))
}
