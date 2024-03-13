package vectors

import "testing"

func remove_elements(nums []int, val int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		if nums[start] == val {
			nums[start], nums[end] = nums[end], nums[start]
			end--
		} else {
			start++
		}
	}

	return end + 1
}

func Test_remove_elements(t *testing.T) {
	tests := []struct {
		nums   []int
		val    int
		result int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for _, test := range tests {
		result := remove_elements(test.nums, test.val)
		if result != test.result {
			panic(test)
		}
	}
}
