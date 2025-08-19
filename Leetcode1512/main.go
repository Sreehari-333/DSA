package main

import "fmt"

func numIdenticalPairs(nums []int) int {

	numOfpairs := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				numOfpairs++
			}
		}
	}
	return numOfpairs
}

func main() {
	// nums := []int{1, 2, 3, 1, 1, 3}
	nums := []int{1, 2, 3}

	fmt.Println(numIdenticalPairs(nums))
}
