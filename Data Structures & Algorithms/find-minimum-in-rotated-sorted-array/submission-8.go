func findMin(nums []int) int {
    res := nums[0]
    left := 0
    right := len(nums)-1
    for left <= right {
        if nums[left] < nums[right] {
            res = min(res, nums[left])
            return res
        }
        mid := left + (right-left)/2
        res = min(res, nums[mid])
        if nums[left] <= nums[mid] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return res
}
