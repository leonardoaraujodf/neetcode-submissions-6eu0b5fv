func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var path []int
	var response [][]int
	var backtrack func(idx, curr int)
	backtrack = func(idx, curr int) {
		if curr == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			response = append(response, tmp)
			return
		}
		
		if idx == len(candidates) {
			return
		}

		prevCandidate := -1 
		for i := idx; i < len(candidates); i++ {
			candidate := candidates[i]
			if prevCandidate == candidate {
				continue
			}

			if curr + candidate <= target {
				path = append(path, candidate)
				backtrack(i + 1, curr + candidate)
				path = path[:len(path)-1] 
			}
			prevCandidate = candidate
		}
	}
	backtrack(0, 0)
	return response
}
