package main

import "fmt"

func maxFreqSum(s string) int {

	freqOfVowels := make(map[rune]int)
	freqOfConsonants := make(map[rune]int)
	maxVowels := 0
	maxConsonants := 0

	for _, str := range s {
		if string(str) == "a" || string(str) == "e" || string(str) == "i" || string(str) == "o" || string(str) == "u" {
			freqOfVowels[str]++
			continue
		}
		freqOfConsonants[str]++

	}

	for _, count := range freqOfVowels {
		if count > maxVowels {
			maxVowels = count
		}
	}

	for _, count := range freqOfConsonants {
		if count > maxConsonants {
			maxConsonants = count
		}
	}

	return maxConsonants + maxVowels
}

func main() {

	s := "successes"
	fmt.Println(maxFreqSum(s))

}
