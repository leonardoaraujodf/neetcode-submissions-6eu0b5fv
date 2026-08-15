func subsetsWithDup(nums []int) [][]int {
	
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)

	var result [][]int
	var subset []int
	var backtrack func(idx int)
	
	backtrack = func(idx int) {
		if idx == len(nums) {
			return
		}
		prev := math.MinInt
		for i := idx; i < len(nums); i++ {
			num := nums[i]
			if prev == num {
				continue
			}

			subset = append(subset, num)		
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			result = append(result, tmp)
			backtrack(i + 1)
			subset = subset[:len(subset)-1]
			prev = num
		}
	}
	
	result = append(result, []int{})
	backtrack(0)
	return result
}
