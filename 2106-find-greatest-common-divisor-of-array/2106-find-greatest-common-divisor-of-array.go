func findGCD(nums []int) int {

	min := 10000
	max := 0
	greatestDiv := 0

	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	for i := 1; i <= max; i++ {
		if min%i == 0 && max%i == 0 {
			greatestDiv = i
		}
	}
	return greatestDiv
}