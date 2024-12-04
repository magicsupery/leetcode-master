package greedy

import "testing"

func maxSubArray(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return -1
	}

	max := nums[0]
	delta := 0
	cur := 0
	for i := 0; i < n; i++ {
		cur = delta + nums[i]
		if cur > max {
			max = cur
		}

		delta += nums[i]
		if delta < 0 {
			delta = 0
		}
	}

	return max

}

func Test_53(t *testing.T) {
	res1 := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if res1 != 6 {
		t.Errorf("error")
	}
}
