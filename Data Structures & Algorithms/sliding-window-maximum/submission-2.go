func maxSlidingWindow(nums []int, k int) []int {
    mq := []int{}
	res := []int{}
	l := 0
	for r := 0; r < len(nums); r++ {
		for len(mq) > 0 && nums[mq[len(mq)-1]] < nums[r] {
			mq = mq[:len(mq)-1]
		}

		mq = append(mq, r)

		if l > mq[0] {
			mq = mq[1:]
		}

		if (r+1) >= k {
			res = append(res, nums[mq[0]])
			l++
		}
	}
	return res
}
