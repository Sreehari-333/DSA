package main

import "fmt"

func maxProfit(prices []int) int {

	left := 0
	right := 1
	profit := 0

	for right < len(prices) {
		if prices[right] > prices[left] {
			if prices[right]-prices[left] > profit {
				profit = prices[right] - prices[left]
			}
		} else {
			left = right
		}
		right++
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
