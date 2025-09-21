package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {

	count := 0
	i := 0

	for i < len(words) {
		for j := 0; j < len(words[i]); j++ {

		}
	}
	return count
}

func main() {

	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	fmt.Println(countConsistentStrings(allowed, words))
}
