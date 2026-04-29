func isValid(s string) bool {
    stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if len(stack) > 0 {
			switch c {
			case ')':
				if stack[len(stack)-1] != '(' {
					return false
				}
			case ']':
				if stack[len(stack)-1] != '[' {
					return false
				}
			case '}':
				if stack[len(stack)-1] != '{' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
