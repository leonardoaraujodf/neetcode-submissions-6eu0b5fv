func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		curr := nums[i]
		if curr > 0 {
			break
		}
		if i > 0 && curr == nums[i-1] {
			continue
		}
		values := twoSum(nums[i+1:], -nums[i])
		// fmt.Println("values: ", values)
		for _, vals := range values {
			res = append(res, []int{curr, vals[0], vals[1]})
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	left := 0
	right := len(nums) - 1
	res := [][]int{}
	// fmt.Println("target: ", target)
	// fmt.Println("nums: ", nums)
	for left < right {
		curr := nums[left] + nums[right]
		// fmt.Printf("left = %d, nums[left] = %d\n", left, nums[left])
		// fmt.Printf("right = %d, nums[right] = %d\n", right, nums[right])
		// fmt.Println("curr: ", curr)
		if curr == target {
			res = append(res, []int{nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
		} else if curr < target {
			left++
		} else if curr > target {
			right--
		}
	}

	return res
}
