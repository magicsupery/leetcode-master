package hash

import (
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0, 10)
	sort.Sort(sort.IntSlice(nums))
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			innerTarget := target - (nums[i] + nums[j])

			left := j + 1
			right := n - 1

			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < n-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				res := nums[left] + nums[right]
				if res == innerTarget {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
				} else if res < innerTarget {
					left++
				} else {
					right--
				}
			}

		}
	}
	return result
}

func Test_18(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	if len(res) != 3 {
		t.Error("fourSum failed")
	}

	res1 := fourSum([]int{2, 2, 2, 2, 2}, 8)
	if len(res1) != 1 {
		t.Error("foursum res1 failed")
	}
}
