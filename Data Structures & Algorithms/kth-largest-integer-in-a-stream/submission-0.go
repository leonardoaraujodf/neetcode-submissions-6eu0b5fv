type PQ []int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	(*pq) = old[:n-1]
	return item
}

type KthLargest struct {
    hp PQ
	k int
}


func Constructor(k int, nums []int) KthLargest {
    var kthLargest KthLargest
	kthLargest.k = k
	kthLargest.hp = PQ(nums)
	heap.Init(&kthLargest.hp)	
	for kthLargest.hp.Len() > k {
		heap.Pop(&kthLargest.hp)
	}
	return kthLargest
}


func (this *KthLargest) Add(val int) int {
	heap.Push(&this.hp, val)
	if this.hp.Len() > this.k {
		heap.Pop(&this.hp)
	}
	return this.hp[0]
}
