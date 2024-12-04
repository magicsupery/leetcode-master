package backtracing

import "testing"

func backtracing(candidates []int, k int, result *[][]int) {

	//ending condition
	if k <= 0 {
		return
	} else if k == 1 {
		for _, candidate := range candidates {
			newResult := []int{candidate}
			*result = append(*result, newResult)
		}
	} else if len(candidates) == k {
		*result = append(*result, candidates)
	} else {
		//backtrace
		resultA := make([][]int, 0)
		pivot := candidates[0]
		backtracing(candidates[1:len(candidates)], k-1, &resultA)
		for i := 0; i < len(resultA); i++ {
			resultA[i] = append(resultA[i], pivot)
		}
		*result = append(*result, resultA...)

		resultB := make([][]int, 0)
		backtracing(candidates[1:len(candidates)], k, &resultB)
		*result = append(*result, resultB...)
	}
}

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || n < k {
		return nil
	}

	result := make([][]int, 0, k)
	candidates := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		candidates = append(candidates, i)
	}
	backtracing(candidates, k, &result)
	return result
}

func Test_77(t *testing.T) {
	res1 := combine(4, 2)
	if len(res1) != 6 {
		t.Errorf("error")
	}

	res2 := combine(1, 1)
	if len(res2) != 1 {
		t.Errorf("error")
	}

	res3 := combine(4, 4)
	if len(res3) != 1 {
		t.Errorf("error")
	}

}
