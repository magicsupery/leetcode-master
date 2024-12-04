package greedy

import "testing"

func jump(nums []int) int {
	n := len(nums)
	atLeastNums := make([]int, n, n)

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j]+j >= i {
				atLeastNum := atLeastNums[j] + 1
				if atLeastNums[i] == 0 || atLeastNum < atLeastNums[i] {
					atLeastNums[i] = atLeastNum
				}
			}
		}
	}
	return atLeastNums[n-1]
}

func jumpSimple(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	barrier := nums[0]
	jumpNum := 1
	start := 1
	for barrier < n-1 {
		step := barrier
		for i := start; i <= step; i++ {
			if i+nums[i] > barrier {
				barrier = i + nums[i]
			}
		}

		jumpNum += 1
	}

	return jumpNum
}

func Test_45(t *testing.T) {
	res1 := jump([]int{2, 3, 1, 1, 4})
	if res1 != 2 {
		t.Errorf("error")
	}

	res2 := jumpSimple([]int{2, 3, 1, 1, 4})
	if res2 != 2 {
		t.Errorf("res2 error")
	}
}
