package main

import "fmt"

// func getFinalState(nums []int, k int, multiplier int) []int {

// 	for k > 0 {
// 		minValue := 10000
// 		minIndex := 0
// 		for i := 0; i < len(nums); i++ {
// 			if nums[i] < minValue {
// 				minValue = nums[i]
// 				minIndex = i
// 			}
// 		}
// 		minValue = minValue * multiplier
// 		nums[minIndex] = minValue
// 		k--
// 	}

// 	return nums
// }

func getFinalState(nums []int, k int, multiplier int) []int {
	for k > 0 {
		minIndex := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
		}
		nums[minIndex] *= multiplier
		k--
	}
	return nums
}

func main() {

	nums := []int{1}
	k := 10
	multiplier := 5
	fmt.Println(getFinalState(nums, k, multiplier))
}
