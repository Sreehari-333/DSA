func reversePrefix(word string, ch byte) string {
	chIndex := -1
	result := ""

	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			chIndex = i
			break
		}
	}

	if chIndex == -1 {
		return word
	}

	for i := chIndex; i >= 0; i-- {
		result += string(word[i])
	}
	result += word[chIndex+1:]
	return result
}