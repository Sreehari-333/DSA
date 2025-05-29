package main

import "fmt"

func isAcronym(words []string, s string) bool {

	for i, word := range words {
		if len(word) == 0 || i >= len(s) || word[0] != s[i] || len(s) > len(words) {
			return false
		}
	}
	return true
}

func main() {

	words := []string{"a", "b", "c"}
	s := "abcd"

	fmt.Println(isAcronym(words, s))
}
