package main

import (
	"fmt"
	"math"
)

func scoreOfString(s string) int {

	sum := 0

	for i := 0; i < len(s)-1; i++ {
		sum += int(math.Abs(float64(s[i]) - float64(s[i+1])))
	}

	return sum
}

func main() {

	s := "zaz"
	fmt.Println(scoreOfString(s))
}
