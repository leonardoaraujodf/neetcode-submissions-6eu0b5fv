func subsets(nums []int) [][]int {
    var result [][]int
    var curr []int
    var backtrack func(startIdx int)
    backtrack = func(startIdx int) {
        tmp := make([]int, len(curr))
        copy(tmp, curr)
        result = append(result, tmp)
        for i := startIdx; i < len(nums); i++ {
            curr = append(curr, nums[i])
            backtrack(i+1)
            curr = curr[:len(curr)-1]
        }
    }
    backtrack(0)
    return result
}
