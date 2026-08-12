func combinationSum(nums []int, target int) [][]int {
    var result [][]int
	var path []int
	var backtrack func(start int, curr int)
	backtrack = func(start int, curr int) {
		if curr == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			num := nums[i]
			if curr + num > target {
				continue
			}
			path = append(path, num)
			backtrack(i, curr + num)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return result
}
