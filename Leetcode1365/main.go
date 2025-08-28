package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {

	arr := []int{}

	for i := 0; i < len(nums); i++ {
		numOfSmall := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				numOfSmall++
			}
		}
		arr = append(arr, numOfSmall)
	}
	return arr
}

func main() {

	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
