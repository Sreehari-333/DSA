package main

import "fmt"

// func minimumOperations(nums []int) int {

// 	numOfOperations := 0
// 	for i := 0; i < len(nums); i++ {
// 		if (nums[i]+1)%3 == 0 || (nums[i]-1)%3 == 0 {
// 			numOfOperations += 1
// 		}
// 		fmt.Println(numOfOperations)
// 	}
// 	return numOfOperations
// }

func minimumOperations(nums []int) int {

	numOfOperations := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%3 != 0 {
			numOfOperations++
		}
	}
	return numOfOperations
}

func main() {
	nums := []int{3, 6, 9}
	fmt.Println(minimumOperations(nums))
}
