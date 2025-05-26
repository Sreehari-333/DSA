package main

import (
	"fmt"
	"strconv"
)

func differenceOfSum(nums []int) int {

	var sumOfDigits int
	var sumOfElements int

	for _, num := range nums {

		strNum := strconv.Itoa(num)

		for _, str := range strNum {
			digit, _ := strconv.Atoi(string(str))

			sumOfDigits += digit
		}

		sumOfElements += num

	}
	return sumOfElements - sumOfDigits
}

func main() {

	num := []int{1, 15, 6, 3}

	fmt.Println(differenceOfSum(num))

}

// var sum int

// 	num := "123"

// 	for _, s := range num {
// 		digit, err := strconv.Atoi(string(s))
// 		if err != nil {
// 			fmt.Println("Error converting:", err)
// 			continue
// 		}
// 		sum += digit
// 		// fmt.Println(s)
// 	}
// 	fmt.Println(sum)
