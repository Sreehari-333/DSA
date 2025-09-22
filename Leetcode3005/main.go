package main

import "fmt"

func maxFrequencyElements(nums []int) int {

	freq := make(map[int]int)
	maxFreq := 0
	sum := 0

	for _, num := range nums {
		freq[num]++
		if freq[num] > maxFreq {
			maxFreq = freq[num]
		}
	}

	for _, val := range freq {
		if val == maxFreq {
			sum += val
		}
	}

	return sum
}

func main() {

	nums := []int{10, 12, 11, 9, 6, 19, 11}
	fmt.Println(maxFrequencyElements(nums))
}
