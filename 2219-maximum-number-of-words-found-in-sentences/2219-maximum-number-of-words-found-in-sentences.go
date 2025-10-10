func mostWordsFound(sentences []string) int {
	maxWord := 0
	for i := 0; i < len(sentences); i++ {
		words := 1
		for _, str := range sentences[i] {
			if string(str) == " " {
				words++
			}
			if maxWord < words {
				maxWord = words
			}
		}
	}
	return maxWord
}