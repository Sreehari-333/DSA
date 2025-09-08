package main

import "fmt"

func largestAltitude(gain []int) int {

	altitude := 0
	heighestAlt := 0

	for i := 0; i < len(gain); i++ {
		altitude += gain[i]

		if altitude > heighestAlt {
			heighestAlt = altitude
		}
	}
	return heighestAlt
}

func main() {

	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}
