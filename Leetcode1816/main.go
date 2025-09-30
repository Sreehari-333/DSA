package main

import "fmt"

func truncateSentence(s string, k int) string {

	return s[:k]
}

func main() {

	s := "Hello how are you Contestant"
	k := 4
	fmt.Println(truncateSentence(s, k))
}
