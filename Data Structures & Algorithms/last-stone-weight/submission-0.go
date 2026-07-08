type PQ []int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
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

func lastStoneWeight(stones []int) int {
    pq := &PQ{}

    for _, stone := range stones {
        *(pq) = append(*pq, stone)
    }

    heap.Init(pq)
    for pq.Len() > 1 {
        x := heap.Pop(pq).(int)
        y := heap.Pop(pq).(int)
        if x == y {
            continue
        } else if x > y {
            x -= y
            heap.Push(pq, x)
        } else {
            y -= x
            heap.Push(pq, y)
        }
    }

    if pq.Len() == 1 {
        return heap.Pop(pq).(int)
    }

    return 0
}