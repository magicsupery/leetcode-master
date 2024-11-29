package str

func doReverse(s []byte, left int, right int) {
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}
}

func reverseStrHelper(s []byte, left int, right int, k int) {
	remain := right - left + 1
	if remain <= k {
		doReverse(s, left, right)
	} else if remain > k && remain <= 2*k {
		doReverse(s, left, left+k-1)
	} else {
		doReverse(s, left, left+k-1)
		reverseStrHelper(s, left+2*k, right, k)
	}
}
func reverseStr(s string, k int) string {
	tmpStr := []byte(s)
	reverseStrHelper(tmpStr, 0, len(s)-1, k)
	return string(tmpStr)
}
