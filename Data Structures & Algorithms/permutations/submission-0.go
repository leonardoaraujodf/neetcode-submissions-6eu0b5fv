func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(curr []int)
	backtrack = func(curr []int) {
		if len(curr) == len(nums) {
			result = append(result, nums)
		}

		for _, num := range nums {
			exists := false
			for _, c := range curr {
				if c == num {
					exists = true
					break
				}
			}
			if exists {
				continue
			}
			curr = append(curr, num)
			backtrack(curr)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack([]int{})
	return result
}
