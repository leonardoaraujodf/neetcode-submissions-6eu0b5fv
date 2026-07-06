// first index refers to length
// second index refers to interval end value
type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i][0] == pq[j][0] {
		return pq[i][1] < pq[j][1]
	} 
	
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([2]int))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0:n-1]
	return item
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

	h := &PriorityQueue{}
	heap.Init(h)
	idx := 0
	res := make([]int, len(queries))
	for _, q := range queries {
        // fmt.Println("q: ", q)
		for idx < len(intervals) && intervals[idx][0] <= q {
			l := intervals[idx][0]
			r := intervals[idx][1]
			heap.Push(h, [2]int{r-l+1, intervals[idx][1]})
			idx++
		}

		for len(*h) > 0 && (*h)[0][1] < q {
            // fmt.Println((*h)[0])
			heap.Pop(h)
		}

        val := -1
        if len(*h) > 0 {
            val = (*h)[0][0]
        }

        for _, idx := range queriesMap[q] {
            res[idx] = val
        }
	}

	return res
}