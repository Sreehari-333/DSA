package main

import "fmt"

func minOperations(nums []int, k int) int {

	var numOfOperation int
	for _, num := range nums {
		if num < k {
			numOfOperation++
		}
	}
	return numOfOperation
}

func main() {

	nums := []int{2, 11, 10, 1, 3}
	k := 10
	fmt.Println(minOperations(nums, k))
}
