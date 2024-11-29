package str

import "testing"

func helper(s string, check string, prefix map[string]bool) bool {
	nxt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == check[nxt] {
			nxt++
			nxt = nxt % len(check)
		} else {
			return false
		}
	}
	return nxt == 0
}
func repeatedSubstringPattern(s string) bool {
	prefixMap := make(map[string]bool)
	for i := 1; i <= len(s)/2; i++ {
		ret := helper(s, s[0:i], prefixMap)
		if ret {
			return ret
		}
	}

	return false
}

func Test_459(t *testing.T) {
	res1 := repeatedSubstringPattern("abab")
	if res1 != true {
		t.Errorf("res1: %t", res1)
	}

	res2 := repeatedSubstringPattern("abcabcabcabc")
	if res2 != true {
		t.Errorf("res2: %t", res2)
	}

	res3 := repeatedSubstringPattern("aba")
	if res3 != false {
		t.Errorf("res3: %t", res3)
	}
}
