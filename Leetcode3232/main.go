package main

import "fmt"

func canAliceWin(nums []int) bool {

	sumOfSingleDig := 0
	sumOfDoubleDig := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 && nums[i] <= 9 {
			sumOfSingleDig += nums[i]
		} else {
			sumOfDoubleDig += nums[i]
		}
	}

	if sumOfSingleDig > sumOfDoubleDig {
		return true
	} else if sumOfDoubleDig > sumOfSingleDig {
		return true
	}

	return false

}

func main() {
	nums := []int{1, 2, 3, 4, 10}
	fmt.Println(canAliceWin(nums))
}
