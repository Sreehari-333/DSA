func capitalizeTitle(title string) string {
	lower := strings.ToLower(title)
	splittedStr := strings.Fields(lower)
	for i := 0; i < len(splittedStr); i++ {
		if len(splittedStr[i]) > 2 {
			splittedStr[i] = strings.Title(splittedStr[i])
		}
	}
	return strings.Join(splittedStr, " ")
}