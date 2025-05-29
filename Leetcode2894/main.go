package main

import "fmt"

func differenceOfSums(n int, m int) int {

	arr := []int{}
	var sumOfDivisible int
	var sumOfNotDivisible int

	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	for _, num := range arr {
		if num%m == 0 {
			sumOfDivisible += num
		} else {
			sumOfNotDivisible += num
		}
	}
	return sumOfNotDivisible - sumOfDivisible

}

func main() {

	fmt.Println(differenceOfSums(10, 3))
}
