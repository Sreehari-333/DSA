package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {

	if len(nums) < 2 {
		return nums
	}

	left := []int{}
	middle := []int{}
	right := []int{}

	for _, val := range nums {
		if val < pivot {
			left = append(left, val)
		} else if val == pivot {
			middle = append(middle, val)
		} else {
			right = append(right, val)
		}
	}

	return append(append(left, middle...), right...)
}

func main() {

	nums := []int{9, 12, 5, 10, 14, 3, 10}
	pivot := 10
	fmt.Println(pivotArray(nums, pivot))
}
