func largestRectangleArea(heights []int) int {
	// index 0 means height
	// index 1 means index
	stack := make([][2]int, 0)
	maxArea := 0
	// to force calculation of all rectangles in the end
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		// strictly increasing stack
		lastIdx := i
		for len(stack) > 0 && heights[i] <= stack[len(stack) - 1][0] {
			lastIdx = stack[len(stack)-1][1]
			height := stack[len(stack)-1][0]
			width := i - lastIdx
			area := height * width
			maxArea = max(maxArea, area)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{heights[i], lastIdx})
		maxArea = max(maxArea, heights[i])
	}

	return maxArea
}
