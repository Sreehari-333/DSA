package main

import (
	"fmt"
	"strconv"
)

func separateDigits(nums []int) []int {

	result := []int{}
	var val string

	for _, num := range nums {
		val = strconv.Itoa(num)
		for i := 0; i < len(val); i++ {
			digit, err := strconv.Atoi(string(val[i]))
			if err == nil {
				result = append(result, digit)
			}
		}
	}
	return result
}

func main() {

	nums := []int{13, 25, 83, 77}
	fmt.Println(separateDigits(nums))
}
