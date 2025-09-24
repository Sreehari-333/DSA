package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {

	result := []int{}
	seen := make(map[int]bool)

	for _, num := range nums {
		seen[num] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !seen[i] {
			result = append(result, i)
		}
	}

	return result
}

func main() {

	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
