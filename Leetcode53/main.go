package main

// import "fmt"

// func maxSubArray(nums []int) int {

// 	currentSum, maxSum := nums[0], nums[0]
// 	// tempStart := 0

// 	for i := 1; i < len(nums); i++ {
// 		if currentSum < 0 {
// 			currentSum = nums[i]
// 			tempStart = i
// 		} else {
// 			currentSum += nums[i]
// 		}
// 		if currentSum > maxSum {
// 			maxSum = currentSum

// 		}
// 	}
// 	return maxSum
// }

// func main() {

// 	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
// 	fmt.Println(maxSubArray(nums))
// }
