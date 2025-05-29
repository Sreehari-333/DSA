package main

import "fmt"

func findWordsContaining(words []string, x byte) []int {

	result := []int{}
	for i, str := range words {
		for j := 0; j < len(str); j++ {
			if str[j] == x {
				result = append(result, i)
				break
			}
		}
	}
	return result
}

func main() {

	words := []string{"leet", "code"}

	var x byte = 'e'

	fmt.Println(findWordsContaining(words, x))
}
