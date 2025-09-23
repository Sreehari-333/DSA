package main

import (
	"fmt"
	// "math"
)

func getCommon(nums1 []int, nums2 []int) int {

	n := min(len(nums1), len(nums2))

	i, j := 0, 0

	for i < n {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1

}

func main() {

	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4}
	fmt.Println(getCommon(nums1, nums2))
}
