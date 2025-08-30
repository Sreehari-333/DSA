package main

import "fmt"

func shuffle(nums []int, n int) []int {

	shuffledArr := []int{}
	for i := 0; i < n; i++ {
		shuffledArr = append(shuffledArr, nums[i], nums[i+n])
	}

	return shuffledArr
}

func main() {

	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3

	fmt.Println(shuffle(nums, n))
}
