package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {

	personDetails := make(map[int]string)

	for i := 0; i < len(names); i++ {
		personDetails[heights[i]] = names[i]
	}

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	for i := 0; i < len(names); i++ {
		names[i] = personDetails[heights[i]]
	}
	return names
}

func main() {

	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	fmt.Println(sortPeople(names, heights))
}
