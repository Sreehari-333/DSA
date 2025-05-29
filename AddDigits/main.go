package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {
	number := num

	for number > 9 {
		sumOfDigits := 0
		strNum := strconv.Itoa(number)
		for _, digit := range strNum {
			n, _ := strconv.Atoi(string(digit))
			sumOfDigits += n
		}
		number = sumOfDigits
	}
	return number
}

func main() {
	fmt.Println(addDigits(38))
}
