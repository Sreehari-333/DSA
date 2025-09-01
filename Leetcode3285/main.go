package main

import "fmt"

func stableMountains(height []int, threshold int) []int {
	numOfStableMountain := []int{}

	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			numOfStableMountain = append(numOfStableMountain, i)
		}
	}
	return numOfStableMountain
}

func main() {

	height := []int{10, 1, 10, 1, 10}
	threshold := 3
	fmt.Println(stableMountains(height, threshold))
}
