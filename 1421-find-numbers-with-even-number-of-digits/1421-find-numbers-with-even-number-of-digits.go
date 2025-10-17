func findNumbers(nums []int) int {
	count := 0

	for _, num := range nums {
		strNum := strconv.Itoa(num)
		if len(strNum)%2 == 0 {
			count++
		}
	}
	return count
}