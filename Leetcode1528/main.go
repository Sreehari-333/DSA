package main

import "fmt"

func restoreString(s string, indices []int) string {

	values := make(map[int]string)
	result := ""

	for i := 0; i < len(indices); i++ {
		values[indices[i]] = string(s[i])
	}

	for i := 0; i < len(indices); i++ {
		result += values[i]

	}

	return result
}

func main() {

	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}

	fmt.Println(restoreString(s, indices))
}
