package main

import "fmt"

func removeElement(nums []int, val int) int {

	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k
}

func main() {
	val := 3
	nums := []int{3, 2, 3, 2}

	fmt.Println(removeElement(nums, val))
}
