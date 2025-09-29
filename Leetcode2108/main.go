package main

import "fmt"

func firstPalindrome(words []string) string {

	for _, str := range words {
		var reversedStr string
		for i := len(str) - 1; i >= 0; i-- {
			reversedStr += string(str[i])
		}

		if str == reversedStr {
			return str
		}
	}
	return ""
}

func main() {

	words := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(firstPalindrome(words))
}
