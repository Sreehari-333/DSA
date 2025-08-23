package main

import "fmt"

func secondLargestNumber(nums []int) int {

	largest := 0
	secondLargest := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > largest {
			secondLargest = largest
			largest = nums[i]
		} else if nums[i] > secondLargest && nums[i] < largest {
			secondLargest = nums[i]
		}
	}
	return secondLargest
}

func main() {
	nums := []int{100, 12, 3, 44, 22, 300}
	fmt.Println(secondLargestNumber(nums))
}
