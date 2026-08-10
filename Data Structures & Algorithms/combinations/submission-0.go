func combine(n int, k int) [][]int {
	var result [][]int
	var curr []int
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(curr) == k {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}

		for i := idx; i <= n; i++ {
			curr = append(curr, i)
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(1)
	return result
}
