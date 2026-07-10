// For this problem we want a min heap.
// That's because we eliminate first the smallest
// elements and keep only the biggest elemenets.
type PriorityQueue []int
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i] < pq[j]
}
func (pq *PriorityQueue) Push(x any) {
    (*pq) = append(*pq, x.(int))
}
func (pq *PriorityQueue) Pop() any {
    tmp := *pq
    n := len(tmp)
    x := tmp[n-1]
    (*pq) = tmp[:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    pq := &PriorityQueue{}
    heap.Init(pq)
    for _, num := range nums {
        heap.Push(pq, num)
    }

    for pq.Len() != k {
        heap.Pop(pq)
    }

    x := heap.Pop(pq)
    return x.(int)
}