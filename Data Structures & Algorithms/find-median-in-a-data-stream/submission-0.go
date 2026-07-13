type MinPriorityQueue []int
func (pq MinPriorityQueue) Len() int { return len(pq) }
func (pq MinPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq MinPriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq *MinPriorityQueue) Push(x any) {
	(*pq) = append(*pq, x.(int))
}
func (pq *MinPriorityQueue) Pop() any {
	tmp := *pq
	n := len(tmp)
	x := tmp[n-1]
	(*pq) = tmp[:n-1]
	return x
}

type MaxPriorityQueue []int
func (pq MaxPriorityQueue) Len() int { return len(pq) }
func (pq MaxPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq MaxPriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}
func (pq *MaxPriorityQueue) Push(x any) {
	(*pq) = append(*pq, x.(int))
}
func (pq *MaxPriorityQueue) Pop() any {
	tmp := *pq
	n := len(tmp)
	x := tmp[n-1]
	(*pq) = tmp[:n-1]
	return x
}


type MedianFinder struct {
    large *MinPriorityQueue
	small *MaxPriorityQueue
}


func Constructor() MedianFinder {
    mf := MedianFinder{
		large: &MinPriorityQueue{},
		small: &MaxPriorityQueue{},
	}
	heap.Init(mf.large)
	heap.Init(mf.small)
	return mf
}


func (this *MedianFinder) AddNum(num int)  {
    large := this.large
	small := this.small
	if large.Len() > 0 && num > (*large)[0] {
		heap.Push(large, num)
	} else {
		heap.Push(small, num)
	}

	// rebalancing
	if large.Len() - small.Len() > 1 {
		x := heap.Pop(large)
		num := x.(int)
		heap.Push(small, num)
	} else if small.Len() - large.Len() > 1 {
		x := heap.Pop(small)
		num := x.(int)
		heap.Push(large, num)
	}
}


func (this *MedianFinder) FindMedian() float64 {
    large := this.large
	small := this.small
	if large.Len() == small.Len() {
		return (float64((*large)[0]) + float64((*small)[0]))/2.0
	} else if large.Len() > small.Len() {
		return float64((*large)[0])
	}
	
	return float64((*small)[0])
}
