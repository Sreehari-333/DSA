package main

import "fmt"

func maxFrequencyElements(nums []int) int {

	seen := make(map[int]bool)
	count := 0

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
		} else {
			count += 2
		}
	}

	if count == 0 {
		return len(nums)
	}
	return count
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maxFrequencyElements(nums))
}
