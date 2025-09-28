package main

import (
	"fmt"
	"strconv"
)

func countDistinctIntegers(nums []int) int {

	seen := make(map[int]bool)

	for _, num := range nums {
		str := strconv.Itoa(num)
		var st string
		for i := len(str) - 1; i >= 0; i-- {
			st += string(str[i])
		}
		val, _ := strconv.Atoi(st)
		nums = append(nums, val)
	}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
		}
	}

	return len(seen)
}

func main() {
	nums := []int{1, 13, 10, 12, 31}
	fmt.Println(countDistinctIntegers(nums))
}
