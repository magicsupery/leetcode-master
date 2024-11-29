package hash

import "testing"

func getNextNumber(n int) int {
	result := 0
	for n >= 10 {
		result += (n % 10) * (n % 10)
		n = n / 10
	}

	result += n * n

	return result
}
func isHappyHelper(n int, resultMap map[int]bool) bool {
	if resultMap[n] {
		return false
	}

	m := getNextNumber(n)
	if m == 1 {
		return true
	}

	resultMap[n] = true
	return isHappyHelper(m, resultMap)

}
func isHappy(n int) bool {
	resultMap := make(map[int]bool)
	return isHappyHelper(n, resultMap)
}

func Test_202(t *testing.T) {
	res := isHappy(19)
	if res != true {
		t.Error("failed")
	}
}
