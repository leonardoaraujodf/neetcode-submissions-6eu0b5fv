type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i int, j int) bool { return pq[i][1] < pq[j][1]}
func (pq PriorityQueue) Swap(i int, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any){
	*pq = append(*pq, x.([2]int))
}
func (pq *PriorityQueue) Pop() any{
	n := len(*pq)
	old := *pq
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
    var pq PriorityQueue
	heap.Init(&pq)
	for key, val := range m {
		heap.Push(&pq, [2]int{key, val})
		if pq.Len() > k {
			heap.Pop(&pq)
		}
	}
    ans := make([]int, pq.Len())
	idx := 0
	for pq.Len() > 0 {
	   item := heap.Pop(&pq)
       ans[idx] = item.([2]int)[0]
	   idx++
	}
    return ans
}
