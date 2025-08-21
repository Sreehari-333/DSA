package main

import "fmt"

func getSneakyNumbers(nums []int) []int {

	m := make(map[int]bool)
	arr := []int{}
	for _, num := range nums {
		if !m[num] {
			m[num] = true
		} else if m[num] {
			arr = append(arr, num)
		}

	}
	return arr
}

func main() {

	nums := []int{0, 1, 0, 1, 2, 3}
	fmt.Println(getSneakyNumbers(nums))
}
