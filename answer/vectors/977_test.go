package vectors

import "testing"

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findMinAbs(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		if intAbs(nums[start]) < intAbs(nums[start+1]) {
			return start
		}

		start++
	}

	return end
}

func sortedSquares(nums []int) []int {
	pivot := findMinAbs(nums)
	start := 0
	result := make([]int, len(nums))
	left, right := pivot, pivot+1

	for left >= 0 && right <= len(nums)-1 {
		if intAbs(nums[left]) < intAbs(nums[right]) {
			result[start] = nums[left] * nums[left]
			left--
		} else {
			result[start] = nums[right] * nums[right]
			right++
		}
		start++
	}

	for left >= 0 {
		result[start] = nums[left] * nums[left]
		left--
		start++
	}

	for right <= len(nums)-1 {
		result[start] = nums[right] * nums[right]
		right++
		start++
	}

	return result
}

func Test_sortedSqured(t *testing.T) {
	tests := []struct {
		nums   []int
		result []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{-1}, []int{1}},
		{[]int{-1, 1}, []int{1, 1}},
		{[]int{-10000, -9999, -7, -5, 0, 0, 10000}, []int{0, 0, 25, 49, 99980001, 100000000, 100000000}},
		{[]int{0, 2}, []int{0, 4}},
		{[]int{}, []int{}},
		{[]int{-2, 0, 3}, []int{0, 4, 9}},
		{[]int{-1, 2, 2}, []int{1, 4, 4}},
		{[]int{-4, -4, 3}, []int{9, 16, 16}},
	}
	for _, test := range tests {
		result := sortedSquares(test.nums)
		for i, v := range result {
			if v != test.result[i] {
				panic(test)
			}
		}
	}
}
