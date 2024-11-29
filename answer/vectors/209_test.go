package vectors

import (
	"math"
	"testing"
)

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	start := 0
	end := len(nums) - 1

	var sum int
	minLengthResult := math.MaxInt
	for start <= end {
		sum = 0

		for i := start; i <= end; i++ {
			sum += nums[i]

			if sum >= target {
				length := i - start + 1
				if length < minLengthResult {
					minLengthResult = length
				}
				break
			}
		}

		start += 1
	}

	if minLengthResult == math.MaxInt {
		return 0
	}

	return minLengthResult
}

func minLengthResultFast(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, 0

	var length int
	minLengthResult := math.MaxInt

	sum := nums[0]
	for (left < len(nums) && right < len(nums)) || left > right {
		if sum >= target {
			length = right - left + 1
			if length < minLengthResult {
				minLengthResult = length
			}

			sum -= nums[left]
			left += 1

		} else {
			right += 1
			if right < len(nums) {
				sum += nums[right]
			}
		}
	}

	if minLengthResult == math.MaxInt {
		return 0
	}

	return minLengthResult
}

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		target int
		nums   []int
		result int
	}{
		{4, []int{1, 4, 4}, 1},
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}
	for _, test := range tests {
		result := minSubArrayLen(test.target, test.nums)
		if result != test.result {
			println("result", result, "test.result", test.result)
			panic(test)
		}
	}
}

func Test_minSubArrayLenFast(t *testing.T) {
	tests := []struct {
		target int
		nums   []int
		result int
	}{
		{4, []int{1, 4, 4}, 1},
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}
	for _, test := range tests {
		result := minLengthResultFast(test.target, test.nums)
		if result != test.result {
			println("result", result, "test.result", test.result)
			panic(test)
		}
	}
}
