package vectors

import (
	"reflect"
	"testing"
)

func getNextPosition(matrix [][]int, i int, j int, direction int) (int, int, int) {
	if direction == 0 {
		if j < len(matrix[0])-1 && matrix[i][j+1] == 0 {
			return i, j + 1, 0
		} else {
			return i + 1, j, 1
		}
	} else if direction == 1 {
		if i < len(matrix)-1 && matrix[i+1][j] == 0 {
			return i + 1, j, 1
		} else {
			return i, j - 1, 2
		}
	} else if direction == 2 {
		if j > 0 && matrix[i][j-1] == 0 {
			return i, j - 1, 2
		} else {
			return i - 1, j, 3
		}
	} else if direction == 3 {
		if i > 0 && matrix[i-1][j] == 0 {
			return i - 1, j, 3
		} else {
			return i, j + 1, 0
		}
	}

	return -1, -1, -1
}

func generateMatrix(n int) [][]int {
	resultMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		resultMatrix[i] = make([]int, n)
	}

	startNum := 1
	i := 0
	j := 0

	//方向控制: 0: 右, 1: 下, 2: 左, 3: 上
	direction := 0

	for startNum <= n*n {
		resultMatrix[i][j] = startNum
		i, j, direction = getNextPosition(resultMatrix, i, j, direction)
		startNum++
	}

	return resultMatrix
}

func Test_generateMatrix(t *testing.T) {
	tests := []struct {
		n      int
		result [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{1, [][]int{{1}}},
	}
	for _, test := range tests {
		result := generateMatrix(test.n)
		if !reflect.DeepEqual(result, test.result) {
			panic(test)
		}
	}
}
