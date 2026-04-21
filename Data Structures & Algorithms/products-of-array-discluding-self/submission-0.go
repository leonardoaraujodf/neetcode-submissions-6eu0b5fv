func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	res := make([]int, len(nums))
	
	prefix[0] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i - 1] * prefix[i - 1]
	}
	// fmt.Println(prefix)
	suffix[n - 1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = nums[i + 1] * suffix[i + 1]
	}
	// fmt.Println(suffix)


	for i := 0; i < n; i++ {
		res[i] = prefix[i] * suffix[i]
	}

	// nums:   [1  2  4  6]
	// prefix: [1  1  2  8]
	// suffix: [48  24  6  1]
	// output: [48 24 12 8]

	return res
}
