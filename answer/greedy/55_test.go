package greedy

import "testing"

func canJump(nums []int) bool {
	n := len(nums)
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	}

	barrier := 0
	for i := 0; i < n-1; i++ {
		if nums[i]+i > barrier {
			barrier = nums[i] + i
		}
		if barrier >= n-1 {
			return true
		} else if barrier <= i {
			return false
		}
	}
	return false
}

func Test_55(t *testing.T) {
	res1 := canJump([]int{2, 3, 1, 1, 4})
	if res1 != true {
		t.Errorf("error")
	}

	res2 := canJump([]int{3, 2, 1, 0, 4})
	if res2 != false {
		t.Errorf("error")
	}
}
