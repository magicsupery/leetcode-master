package stackqueue

import (
	"strconv"
	"testing"
)

func isSign(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}

	return false
}

func cal(a string, b string, sign string) int {
	aNum, _ := strconv.Atoi(a)
	bNum, _ := strconv.Atoi(b)

	if sign == "+" {
		return aNum + bNum
	} else if sign == "-" {
		return aNum - bNum
	} else if sign == "*" {
		return aNum * bNum
	} else if sign == "/" {
		return aNum / bNum
	}

	return 0
}
func evalRPN(tokens []string) int {
	stack := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if !isSign(token) {
			stack = append(stack, token)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			c := cal(a, b, token)
			stack = append(stack, strconv.Itoa(c))
		}
	}

	result, _ := strconv.Atoi(stack[0])
	return result
}

func Test_1047(t *testing.T) {
	res1 := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	if res1 != 22 {
		t.Error("Test case 1 failed")
	}
}
