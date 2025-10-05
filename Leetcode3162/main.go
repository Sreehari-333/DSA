package main

import "fmt"

func numberOfPairs(nums1 []int, nums2 []int, k int) int {

	i := 0
	numOfGoodPairs := 0
	for i < len(nums1) {
		j := 0
		for j < len(nums2) {
			if nums1[i]%(nums2[j]*k) == 0 {
				numOfGoodPairs++
			}

			j++
		}
		i++
	}
	return numOfGoodPairs
}

func main() {

	nums1 := []int{1, 2, 4, 12}
	nums2 := []int{2, 4}
	k := 3
	fmt.Println(numberOfPairs(nums1, nums2, k))
}
