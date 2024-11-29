package hash

func intersection(nums1 []int, nums2 []int) []int {
	num1Map := make(map[int]bool)
	for _, num := range nums1 {
		num1Map[num] = true
	}

	result := make([]int, 0, len(nums1))
	for _, num := range nums2 {
		if _, ok := num1Map[num]; ok {
			result = append(result, num)
			delete(num1Map, num)
		}
	}

	return result
}
