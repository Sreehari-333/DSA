package main

import "fmt"

func subsetXORSum(nums []int) int {
	sumTotal := 0

	for _, num := range nums {
		sumTotal |= num
	}
	return sumTotal << (len(nums) - 1)
}

func main() {

	nums := []int{1, 3}
	fmt.Println(subsetXORSum(nums))
}
