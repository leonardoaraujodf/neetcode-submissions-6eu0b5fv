type info struct {
	temperature int
	index int
}

func dailyTemperatures(temperatures []int) []int {
	stack := []info{}
	res := make([]int, len(temperatures))
	for idx, temp := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].temperature < temp {
			elem := stack[len(stack)-1]
			res[elem.index] = idx - elem.index
			stack = stack[:len(stack)-1]
		}
		newElem := info{temp, idx}
		stack = append(stack, newElem)
	}
	return res
}
