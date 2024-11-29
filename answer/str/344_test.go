package str

func reverseString(s []byte) {
	n := len(s)
	left := 0
	right := n - 1

	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}
}
