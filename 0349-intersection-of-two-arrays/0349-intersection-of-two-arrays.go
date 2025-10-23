func intersection(nums1 []int, nums2 []int) []int {

	intersect := make(map[int]bool)
	result := []int{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				intersect[nums1[i]] = true
			}
		}
	}

	for key := range intersect {
		result = append(result, key)
	}
	return result
}