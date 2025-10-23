func intersection(nums1 []int, nums2 []int) []int {

	intersect := make(map[int]bool)
	result := []int{}

	for _, num := range nums1 {
		intersect[num] = true
	}

	for _, num := range nums2 {

		if intersect[num] {
			result = append(result, num)
			intersect[num] = false
		}
	}
	return result
}