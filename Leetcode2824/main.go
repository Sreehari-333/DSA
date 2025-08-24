package main

import "fmt"

func countPairs(nums []int, target int) int {

	numOfPairs := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				numOfPairs++
			}
		}
	}
	return numOfPairs
}

func main() {

	nums := []int{-1, 1, 2, 3, 1}
	target := 2
	fmt.Println(countPairs(nums, target))
}
