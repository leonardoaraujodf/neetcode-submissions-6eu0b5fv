func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	res := nums[0]
	for left <= right {
		if nums[left] < nums[right] {
			if nums[left] < res {
				res = nums[left]
			}
			break
		}
		mid := left + (right - left)/2
		if nums[mid] < res {
			res = nums[mid]
		}

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
