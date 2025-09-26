package main

import "fmt"

func uniqueOccurrences(arr []int) bool {

	freq := make(map[int]int)
	seen := make(map[int]bool)

	for _, val := range arr {
		freq[val]++
	}

	for _, val := range freq {
		if !seen[val] {
			seen[val] = true
		} else {
			return false
		}
	}

	return true

}

func main() {

	arr := []int{1, 2, 2, 1, 1, 3, 4}
	fmt.Println(uniqueOccurrences(arr))

}
