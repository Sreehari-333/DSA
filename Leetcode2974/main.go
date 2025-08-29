package main

import (
	"fmt"
	"sort"
)

func numberGame(nums []int) []int {

	resultArr := []int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {
		if i+1 < len(nums) {
			resultArr = append(resultArr, nums[i+1], nums[i])
		}
	}

	return resultArr
}

func main() {

	nums := []int{5, 4, 2, 3}
	fmt.Println(numberGame(nums))
}
