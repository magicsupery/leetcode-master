package vectors

import "testing"

func search_liner(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}

func search_binary_help(nums []int, left int, right int, target int) int {
	if left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			return search_binary_help(nums, mid+1, right, target)
		} else {
			return search_binary_help(nums, left, mid-1, target)
		}
	} else {
		return -1
	}
}

func search_binary(nums []int, target int) int {
	return search_binary_help(nums, 0, len(nums)-1, target)
}

func search_binary_loop(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func Test_LinerSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, -5, -1},
	}

	for _, test := range tests {
		if got := search_liner(test.nums, test.target); got != test.want {
			t.Errorf("search_liner(%v, %v) = %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}

func Test_BinarySearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, -5, -1},
	}

	for _, test := range tests {
		if got := search_binary(test.nums, test.target); got != test.want {
			t.Errorf("search_binary(%v, %v) = %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}

func Test_BinarySearchLoop(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, -5, -1},
	}

	for _, test := range tests {
		if got := search_binary_loop(test.nums, test.target); got != test.want {
			t.Errorf("search_binary_loop(%v, %v) = %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}
