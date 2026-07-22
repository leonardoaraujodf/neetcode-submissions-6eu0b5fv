// The first index refers to length
// The second index refers to end time
type PriorityQueue [][2]int
func (pq PriorityQueue) Len() int {
    return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
    if pq[i][0] == pq[j][0] {
        return pq[i][1] < pq[j][1]
    }

    return pq[i][0] < pq[j][0]
}
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}
func(pq *PriorityQueue) Push(x any) {
    (*pq) = append(*pq, x.([2]int))
}
func (pq *PriorityQueue) Pop() any {
    tmp := *pq
    n := len(tmp)
    x := tmp[n-1]
    (*pq) = tmp[:n-1]
    return x
}

func minInterval(intervals [][]int, queries []int) []int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    queriesMap := map[int][]int{}
    for idx, val := range queries {
        queriesMap[val] = append(queriesMap[val], idx)
    }
    sort.Slice(queries, func(i, j int) bool {
        return queries[i] < queries[j]
    })

    res := make([]int, len(queries))
    idx := 0
    pq := &PriorityQueue{}
    heap.Init(pq)
    for _, q := range queries {
        for idx < len(intervals) && intervals[idx][0] <= q {
            interval := intervals[idx]
            l, r := interval[0], interval[1]
            length := r - l + 1
            heap.Push(pq, [2]int{length, r})
            idx++
        }

        val := -1
        for pq.Len() > 0 && (*pq)[0][1] < q {
            heap.Pop(pq)
        }
        if pq.Len() > 0 {
            val = (*pq)[0][0]
        }

        for _, idx := range queriesMap[q] {
            res[idx] = val
        }
    }
    return res
}