package main

import "fmt"

func finalPrices(prices []int) []int {
	n := len(prices)
	answer := make([]int, n)
	for i := 0; i < n; i++ {
		discount := 0
		for j := i + 1; j < n; j++ {
			if prices[j] <= prices[i] {
				discount = prices[j]
				break
			}
		}
		answer[i] = prices[i] - discount
	}
	return answer
}

func main() {

	prices := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(prices))
}
