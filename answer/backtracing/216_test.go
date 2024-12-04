package backtracing

import "testing"

func isValid(start int, k int, n int) bool {
	total := 0
	for i := 0; i < k; i++ {
		total += (start + i)
		if total > n {
			return false
		}
	}
	return true
}
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0, 10)
	end := 9

	var traceback func(start int, k int, n int, path []int)
	traceback = func(start int, k int, n int, path []int) {
		if n <= 0 {
			return
		}

		//last one
		if k == 1 {
			if start <= n && n <= end {
				finalPath := make([]int, len(path)+1)
				copy(finalPath, path)
				finalPath[len(path)] = n
				result = append(result, finalPath)
			}
		} else {
			for i := start; i <= end; i++ {
				path = append(path, i)
				if isValid(i+1, k-1, n-i) {
					traceback(i+1, k-1, n-i, path)
					path = path[:len(path)-1]

				}
			}
		}

	}

	traceback(1, k, n, []int{})
	return result
}

func Test_216(t *testing.T) {
	res1 := combinationSum3(3, 7)
	if res1[0][0] != 1 || res1[0][1] != 2 || res1[0][2] != 4 {
		t.Errorf("error")
	}

	res2 := combinationSum3(8, 36)
	if res2[0][0] != 1 || res2[0][1] != 2 || res2[0][2] != 3 || res2[0][3] != 4 || res2[0][4] != 5 || res2[0][5] != 6 || res2[0][6] != 7 || res2[0][7] != 8 {
		t.Errorf("error")
	}
}
