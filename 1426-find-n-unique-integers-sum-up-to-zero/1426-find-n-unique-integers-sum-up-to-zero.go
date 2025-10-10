func sumZero(n int) []int {

	var result = []int{}

	for i := 1; i < n/2+1; i++ {
		result = append(result, i, -i)
	}

	if n%2 != 0 {
		result = append(result, 0)
	}

	return result
}