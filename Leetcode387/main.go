package main

import "fmt"

func firstUniqChar(s string) int {

	seen := make(map[rune]int)

	for _, val := range s {
		seen[val]++
	}

	for i, val := range s {
		if seen[val] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
