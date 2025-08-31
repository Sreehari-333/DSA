package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	mergedStr := ""
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		mergedStr += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		mergedStr += string(word1[i])
		i++
	}

	for j < len(word2) {
		mergedStr += string(word1[j])
		j++
	}

	return mergedStr
}

func main() {
	word1 := "abcd"
	word2 := "pq"
	fmt.Println(mergeAlternately(word1, word2))
}
