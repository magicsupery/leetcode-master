package hash

func fourSumCount_complex(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	result := 0
	n := len(nums1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					if nums1[i]+nums2[j]+nums3[k]+nums4[l] == 0 {
						result += 1
					}
				}
			}
		}
	}

	return result
}

func fourSumCount_hash(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)

	result1Map := make(map[int]uint)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result1Map[nums1[i]+nums2[j]] += 1
		}
	}

	result2Map := make(map[int]uint)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result2Map[nums3[i]+nums4[j]] += 1
		}
	}

	var count uint
	count = 0
	for k, v := range result1Map {
		count += result2Map[-k] * v
	}

	return int(count)
}
