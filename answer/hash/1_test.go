package hash

import (
	"testing"
)

func twoSum_complex(nums []int, target int) []int {
	var i, j int

	n := len(nums)
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{i, j}
}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok {
			return []int{j, i}
		}
		numsMap[num] = i
	}

	return nil
}

func Test_2(t *testing.T) {
	res1 := twoSum_complex([]int{2, 7, 11, 15}, 9)
	if res1 == nil || res1[0] != 0 || res1[1] != 1 {
		t.Error("twoSum_complex failed")
	}

	res2 := twoSum([]int{2, 7, 11, 15}, 9)
	if res2 == nil || res2[0] != 0 || res2[1] != 1 {
		t.Error("twoSum failed")
	}
}
