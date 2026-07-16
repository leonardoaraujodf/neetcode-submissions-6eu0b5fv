func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    res := nums[0]
    for left <= right {
        if nums[left] < nums[right] {
            res = min(res, nums[left])
            return res
        }
        mid := left + (right - left)/2
        res = min(res, nums[mid])
        if nums[left] <= nums[mid] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return res
}
