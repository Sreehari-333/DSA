package main

import "fmt"

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1

	result := 0

	for left < right {

		heightOfContainer := 0
		if height[left] < height[right] {
			heightOfContainer = height[left]
		} else {
			heightOfContainer = height[right]
		}

		width := right - left
		area := width * heightOfContainer
		if area > result {
			result = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result

}

func main() {

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(height))
}
