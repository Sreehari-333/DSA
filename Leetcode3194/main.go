package main

import (
	"fmt"
	"sort"
)

func minimumAverage(nums []int) float64 {
	var min float64
	n := len(nums)
	sort.Ints(nums)
	min = float64(nums[0]+nums[n-1]) / 2.
	for i := 1; i < n/2; i++ {
		if float64(nums[i]+nums[n-1-i])/2. < min {
			min = float64(nums[i]+nums[n-1-i]) / 2.
		}
	}
	return min
}

func main() {

	nums := []int{7, 8, 3, 4, 15, 13, 4, 1}
	fmt.Println(minimumAverage(nums))
}
