func findMin(nums []int) int {
    res := nums[0]
    left, right := 0, len(nums)-1
    for left <= right {
        if nums[left] < nums[right] {
            res = min(res, nums[left])
            return res
        }
        
        mid := left + (right - left)/2
        val := nums[mid]
        res = min(res, val)
        if val >= nums[left] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return res
}
