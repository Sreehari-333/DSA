package main

import "fmt"

func singleNumber(nums []int) int {

	freq := make(map[int]int)
	var result int

	for _, num := range nums {
		freq[num]++
	}

	for val, count := range freq {
		if count == 1 {
			result = val
		}
	}
	return result
}

func main() {

	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
