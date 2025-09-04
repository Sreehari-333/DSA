package main

import "fmt"

func sumOfSquares(nums []int) int {

	sumOfElements := 0
	n := len(nums)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sumOfElements += (nums[i-1] * nums[i-1])
		}
	}
	return sumOfElements
}

func main() {
	nums := []int{2, 7, 1, 19, 18, 3}
	fmt.Println(sumOfSquares(nums))
}
