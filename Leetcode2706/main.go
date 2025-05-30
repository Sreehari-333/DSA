package main

import (
	"fmt"
	"sort"
)

func buyChoco(prices []int, money int) int {

	sort.Ints(prices)
	i := 0

	if prices[i]+prices[i+1] <= money {
		return money - (prices[i] + prices[i+1])
	}
	return money
}

func main() {

	prices := []int{3, 2, 3}
	money := 3

	fmt.Println(buyChoco(prices, money))
}
