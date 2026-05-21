func search(nums []int, target int) int {
	pivotIdx := findMinIdx(nums)
	// fmt.Println(pivotIdx)
	if nums[pivotIdx] == target {
		return pivotIdx
	}

	result := binarySearch(0, pivotIdx-1, nums, target)
	if result != -1 {
		return result
	}

	return binarySearch(pivotIdx, len(nums)-1, nums, target)
}

func binarySearch(left int, right int, nums []int, target int) int {
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func findMinIdx(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return left
		}
		mid := left + (right - left)/2
		if nums[mid] < nums[left] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}