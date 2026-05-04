func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		// fmt.Printf("token %s\n", token)
		num, err := strconv.Atoi(token)
		if err == nil {
			//fmt.Println(num)
			// means it is a number
			stack = append(stack, num)
			continue
		}

		var operation func(a int, b int) (int)
		switch token {
		case "+":
			operation = func(a int, b int)(int) {
				return a + b
			}
			// fmt.Println("Here")
		case "-":
			operation = func(a int, b int)(int) {
				return a - b
			}
		case "*":
			operation = func(a int, b int)(int) {
				return a * b
			}
		case "/":
			operation = func(a int, b int)(int) {
				return a / b
			}
		}

		a := stack[len(stack)-2]
		b := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		result := operation(a, b)
		stack = append(stack, result)
	}
	return stack[len(stack)-1]
}
