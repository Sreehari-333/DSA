func reverseVowels(s string) string {

	vowels := ""
	reversedVowels := ""
	j := 0
	result := ""

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" || string(s[i]) == "E" || string(s[i]) == "I" || string(s[i]) == "O" || string(s[i]) == "U" || string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" {
			vowels += string(s[i])
		}
	}

	for i := len(vowels) - 1; i >= 0; i-- {
		reversedVowels += string(vowels[i])
	}

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" || string(s[i]) == "E" || string(s[i]) == "I" || string(s[i]) == "O" || string(s[i]) == "U" || string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" {
			result += string(reversedVowels[j])
			j++
		} else {
			result += string(s[i])
		}
	}

	return result
}