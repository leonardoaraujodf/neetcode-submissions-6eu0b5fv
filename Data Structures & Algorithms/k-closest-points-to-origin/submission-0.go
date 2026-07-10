type Point struct {
    distance float64
    idx int
}

type PriorityQueue []Point

func (pq PriorityQueue) Len() int {
    return len(pq)
}

// we want a max heap. This is because we would like the 
// k closest points to the origin, which translates to the
// points that are least distant from origin. We want to remove
// the bigger distances first.
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].distance > pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
    (*pq) = append(*pq, x.(Point))
}

func (pq *PriorityQueue) Pop() any {
    tmp := *pq
    n := len(tmp)
    x := tmp[n-1]
    (*pq) = tmp[:n-1]
    return x
}

func kClosest(points [][]int, k int) [][]int {
    pq := &PriorityQueue{}
    heap.Init(pq)
    for idx, point := range points {
        x := point[0]
        y := point[1]
        p := Point{
            distance: math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2)),
            idx: idx,
        }
        heap.Push(pq, p)
    }
    for pq.Len() > k {
        heap.Pop(pq)
    }

    ans := [][]int{}
    for pq.Len() > 0 {
        item := heap.Pop(pq)
        p := item.(Point)
        ans = append(ans, points[p.idx])
    }
    return ans
}