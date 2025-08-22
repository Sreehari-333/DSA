package main

import "fmt"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {

	countOfEmployee := 0
	for _, hour := range hours {
		if hour >= target {
			countOfEmployee++
		}
	}
	return countOfEmployee
}
func main() {
	hours := []int{0, 1, 2, 3, 4}
	target := 2

	fmt.Println(numberOfEmployeesWhoMetTarget(hours, target))
}
