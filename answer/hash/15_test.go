package hash

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	var result [][]int
	resultMap := make(map[[3]int]bool)
	n := len(nums)

	for i := 0; i < n; i++ {
		left := i + 1
		right := n - 1

		target := -nums[i]
		for left < right {
			//if left == i {
			//	left++
			//	continue
			//}
			//
			//if right == i {
			//	right--
			//	continue
			//}

			res := nums[left] + nums[right]
			if res == target {
				resultMap[[3]int{nums[i], nums[left], nums[right]}] = true
				left++
			} else if res < target {
				left++
			} else {
				right--
			}
		}
	}

	for k, _ := range resultMap {
		result = append(result, []int{k[0], k[1], k[2]})
	}
	return result
}

func Test_15(t *testing.T) {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	if len(res) != 2 {
		t.Error("threeSum failed")
	}

	res1 := threeSum([]int{1, 2, -2, -1})
	if len(res1) != 0 {
		t.Error("threeSum failed")
	}
}
