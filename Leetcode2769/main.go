package main

import "fmt"

func theMaximumAchievableX(num int, t int) int {
	return num + (2 * t)
}

func main() {
	num := 3
	t := 2
	fmt.Println(theMaximumAchievableX(num, t))
}
