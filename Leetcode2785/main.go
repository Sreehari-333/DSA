package main

import (
	"fmt"
	"sort"
)

func sortVowels(s string) string {

	askiVal := []int{}
	index := []int{}

	for i := 0; i < len(s); i++ {

		if string(s[i]) == "A" || string(s[i]) == "E" || string(s[i]) == "I" || string(s[i]) == "O" || string(s[i]) == "U" || string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" {
			askiVal = append(askiVal, int(s[i]))
			index = append(index, i)
		}

	}

	sort.Ints(askiVal)
	b := []byte(s)
	j := 0
	for i := 0; i < len(index); i++ {
		b[index[i]] = byte(askiVal[j])
		j++
	}

	return string(b)

}

func main() {

	s := "lEetcOde"
	fmt.Println(sortVowels(s))
}
