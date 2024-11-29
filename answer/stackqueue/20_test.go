package stackqueue

func isValid(s string) bool {
	stack := make([]int32, 0, len(s))
	for _, char := range s {
		if char == '[' || char == '{' || char == '(' {
			stack = append(stack, char)
		} else if char == ']' {
			if len(stack) <= 0 {
				return false
			}
			check := stack[len(stack)-1]
			if check != '[' {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else if char == '}' {
			if len(stack) <= 0 {
				return false
			}
			check := stack[len(stack)-1]
			if check != '{' {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else if char == ')' {
			if len(stack) <= 0 {
				return false
			}
			check := stack[len(stack)-1]
			if check != '(' {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
