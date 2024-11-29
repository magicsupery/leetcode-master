package stackqueue

import (
	"sort"
	"testing"
)

func swapLeast(container []int, count map[int]int, element int) {
	if count[container[0]] >= count[element] {
		return
	} else {
		container[0] = element
		sort.Slice(container, func(i, j int) bool {
			return count[container[i]] < count[container[j]]
		})
	}

}
func replaceLeast(container []int, count map[int]int, element int, k int) []int {
	n := len(container)
	if n < k-1 {
		container = append(container, element)
	} else if n == k-1 {
		container = append(container, element)
		sort.Slice(container, func(i, j int) bool {
			return count[container[i]] < count[container[j]]
		})
	} else {
		swapLeast(container, count, element)
	}

	return container
}
func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num] += 1
	}

	result := make([]int, 0, k)
	for num, _ := range numsMap {
		result = replaceLeast(result, numsMap, num, k)
	}

	return result
}

func Test_347(t *testing.T) {
}
