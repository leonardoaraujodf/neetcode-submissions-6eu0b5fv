func twoSum(nums []int, target int) []int {
    // key is the value in nums slice
    // value is the index in nums slice
    m := make(map[int]int)
    for idx, val := range nums {
        idx2, exists := m[target - val]
        if exists {
            return []int{idx2, idx}
        }
        m[val] = idx
    }
    return []int{}
}
