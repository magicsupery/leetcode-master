package str

import "testing"

func strStr(haystack string, needle string) int {
	nxt := 0
	for i := 0; i < len(haystack); {
		if haystack[i] == needle[nxt] {
			if nxt == len(needle)-1 {
				return i - nxt
			}

			nxt++
			i++
		} else {
			i = i - nxt + 1
			nxt = 0

		}
	}

	return -1
}

func Test_28(t *testing.T) {
	res1 := strStr("hello", "ll")
	if res1 != 2 {
		t.Errorf("res1: %d", res1)
	}

	res2 := strStr("mississippi", "issip")
	if res2 != 4 {
		t.Errorf("res2: %d", res2)
	}
}
