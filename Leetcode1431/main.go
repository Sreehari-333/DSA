package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {

	var arr []bool
	max := 0

	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if (candies[i] + extraCandies) >= max {
			arr = append(arr, true)
		} else {
			arr = append(arr, false)
		}
	}
	return arr
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	fmt.Println(kidsWithCandies(candies, extraCandies))
}
