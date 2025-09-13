package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {

	targets := []int{}

	sort.Ints(nums)

	for i, num := range nums {
		if num == target {
			targets = append(targets, i)
		}
	}
	return targets
}

func main() {

	nums := []int{1, 2, 5, 2, 3}
	target := 2

	fmt.Println(targetIndices(nums, target))
}
