package main

import (
	"fmt"
	"sort"
)

func findDisappearedNumbers(nums []int) []int {

	sort.Ints(nums)
	result := []int{}
	i := nums[0]

	for i < len(nums) {

	}
}

func main() {

	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
