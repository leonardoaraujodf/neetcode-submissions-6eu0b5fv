func subsets(nums []int) [][]int {
    var result [][]int
    var backtrack func(curr []int, currIdx int)
    backtrack = func(curr []int, currIdx int) {
        tmp := make([]int, len(curr))
        copy(tmp, curr)
        result = append(result, tmp)
        for _, num := range nums {
            if num <= currIdx {
                continue
            }
            curr = append(curr, num)
            backtrack(curr, num)
            curr = curr[:len(curr)-1]
        }
    }
    backtrack([]int{}, math.MinInt)
    return result
}
