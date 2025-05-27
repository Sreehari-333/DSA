package main

import (
	"fmt"
	"strconv"
)

func countEven(num int) int {

	var sumOfDigits int
	var digit int
	result := []int{}

	for i := 2; i <= num; i++ {

		strNum := strconv.Itoa(i)
		sumOfDigits = 0

		for _, str := range strNum {
			digit, _ = strconv.Atoi(string(str))
			sumOfDigits += digit
		}
		if sumOfDigits%2 == 0 {
			result = append(result, i)
		}
	}

	return len(result)
}

func main() {
	fmt.Println(countEven(30))
}
