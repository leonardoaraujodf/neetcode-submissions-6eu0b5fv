func hasDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, val := range nums {
        _, exists := m[val]
        if exists {
            return true
        }
        m[val] = true
    }
    return false
}
