package main

import "fmt"

func sumOfUnique(nums []int) int {

	sumOfUnique := 0
	freq := make(map[int]int, len(nums))

	for _, num := range nums {
		freq[num]++
	}

	for num, count := range freq {
		if count == 1 {
			sumOfUnique += num
		}
	}

	return sumOfUnique
}

func main() {

	nums := []int{1, 2, 3, 2}

	fmt.Println(sumOfUnique(nums))
}
