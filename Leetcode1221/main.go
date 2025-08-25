package main

import "fmt"

func balancedStringSplit(s string) int {

	numOfL := 0
	numOfR := 0
	numOfSubStr := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "L" {
			numOfL++
		} else if string(s[i]) == "R" {
			numOfR++
		}

		if numOfL == numOfR {
			numOfSubStr++
		}
	}

	return numOfSubStr
}

func main() {
	s := "RLRRLLRLRL"

	fmt.Println(balancedStringSplit(s))
}
