type MonotonicQueue struct {
	maxq list.List
}

func (mq *MonotonicQueue) Push(elem int) {
	for mq.maxq.Len() > 0 && mq.maxq.Back().Value.(int) < elem {
		mq.maxq.Remove(mq.maxq.Back())
	}
	mq.maxq.PushBack(elem)
}

func (mq *MonotonicQueue) Max() int {
	return mq.maxq.Front().Value.(int)
}

func (mq *MonotonicQueue) Pop(elem int) {
	if elem == mq.maxq.Front().Value.(int) {
		mq.maxq.Remove(mq.maxq.Front())
	}
}

func maxSlidingWindow(nums []int, k int) []int {
    mq := MonotonicQueue{}
	for i := 0; i < k; i++ {
		mq.Push(nums[i])
	}

	res := []int{mq.Max()}
	for i := k; i < len(nums); i++ {
		mq.Push(nums[i])
		mq.Pop(nums[i-k])
		res = append(res, mq.Max())
	}
	return res
}
