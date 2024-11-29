package stackqueue

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	result := make([]int, 0, n-k+1)

	// init the heap
	h := IntHeap(nums[0 : k-1])
	heap.Init(&h)
	heapMax := k
	// slide windows
	for i := 0; i < n-k+1; i++ {

	}
	return result
}
