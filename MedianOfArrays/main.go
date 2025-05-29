package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	mergedArray := append(nums1, nums2...)

	sort.Ints(mergedArray)

	if len(mergedArray)%2 != 0 {
		mid := (len(mergedArray) - 1) / 2
		return float64(mergedArray[mid])
	}
	mid2 := mergedArray[len(mergedArray)/2]
	mid1 := mergedArray[(len(mergedArray)/2)-1]

	fmt.Println(mid2)
	fmt.Println(mid1)

	result := (float64(mid1) + float64(mid2)) / 2
	fmt.Println(result)

	return float64(result)

}

func main() {
	arr1 := []int{1, 2}
	arr2 := []int{3, 4}
	fmt.Printf("%f", findMedianSortedArrays(arr1, arr2))
}
