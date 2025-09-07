package main

import (
	"fmt"
	"sort"
)

func main() {

	num1 := []int{10}
	num2 := []int{5}

	fmt.Println(addedInteger(num1, num2))
}

func addedInteger(nums1 []int, nums2 []int) int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	x := nums2[0] - nums1[0]

	return x
}
