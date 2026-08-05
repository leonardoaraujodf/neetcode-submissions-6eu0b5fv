func generateParenthesis(n int) []string {
	var result []string
	var stack []byte
	var backtrack func(opened int, closed int)
	backtrack = func(opened int, closed int) {
		if opened == n && closed == n {
			result = append(result, string(stack))
			return
		}
		// try open a parenthesis
		if opened < n {
			stack = append(stack, '(')
			backtrack(opened+1, closed)
			stack = stack[:len(stack)-1]

		}
		if closed < opened {
			stack = append(stack, ')')
			backtrack(opened, closed+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return result
}
