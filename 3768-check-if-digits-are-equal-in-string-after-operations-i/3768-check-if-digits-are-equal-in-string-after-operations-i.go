func hasSameDigits(s string) bool {

	for len(s) > 2 {
		newStr := ""

		for i := 0; i < len(s)-1; i++ {
			a := int(s[i] - '0')
			b := int(s[i+1] - '0')
			sum := (a + b) % 10
			newStr += strconv.Itoa(sum)
		}

		s = newStr
	}

	return s[0] == s[1]
}