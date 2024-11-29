package str

import (
	"fmt"
	"testing"
)

func reverseWords(s string) string {
	tmpBytes := make([][]byte, 0, 1)

	n := len(s)
	left := -1
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			if left != -1 {
				tmpByte := make([]byte, 0, i-left)
				for j := left; j < i; j++ {
					tmpByte = append(tmpByte, s[j])
				}
				tmpBytes = append(tmpBytes, tmpByte)

				left = -1
			}
		} else if left == -1 {
			left = i
		}
	}

	if left != -1 {
		tmpByte := make([]byte, 0, n-left)
		for j := left; j < n; j++ {
			tmpByte = append(tmpByte, s[j])
		}
		tmpBytes = append(tmpBytes, tmpByte)
	}

	result := ""
	for i := len(tmpBytes) - 1; i >= 0; i-- {
		result += string(tmpBytes[i])
		if i != 0 {
			result += " "
		}
	}

	return result
}

func Test_151(t *testing.T) {
	res1 := reverseWords("the sky is blue")
	if res1 != "blue is sky the" {
		fmt.Println("res1:", res1)
		t.Error("151 failed")
	}

	res2 := reverseWords("  hello world!  ")
	if res2 != "world! hello" {
		t.Error("151 failed")
	}
}
