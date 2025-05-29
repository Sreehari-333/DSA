package main

import "fmt"

func averageValue(nums []int) int {

	var result = []int{}
	var result1 int
	for _, num := range nums {
		if num%2 == 0 && num%3 == 0 {
			result = append(result, num)
		}
	}

	if len(result) == 0 {
		return 0
	} else if len(result) == 1 {
		return result[0]
	} else {
		for _, num := range result {
			result1 += num
		}
		return result1 / len(result)
	}

}

func main() {

	nums := []int{43, 9, 75, 76, 25, 96, 46, 85, 19, 29, 88, 2, 5, 24, 60, 26, 76, 24, 96, 82, 97, 97, 72, 35, 21, 77, 82, 30, 94, 55, 76, 94, 51}

	fmt.Println(averageValue(nums))
}
