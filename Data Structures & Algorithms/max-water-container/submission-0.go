func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	res := 0
	for left < right {
		length := right - left
		height := min(heights[left], heights[right])
		res = max(res, length * height)
		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}
	return res
}
