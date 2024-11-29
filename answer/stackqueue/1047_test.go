package stackqueue

func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			check := stack[len(stack)-1]
			if check == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
