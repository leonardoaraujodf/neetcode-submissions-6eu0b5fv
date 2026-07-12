type Task struct {
    count int
    time int
}
type PriorityQueue []Task

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].count > pq[j].count
}
func (pq *PriorityQueue) Push(x any) {
    (*pq) = append(*pq, x.(Task))
}
func (pq *PriorityQueue) Pop() any {
    tmp := *pq
    n := len(tmp)
    x := tmp[n-1]
    (*pq) = tmp[:n-1]
    return x
}

func leastInterval(tasks []byte, n int) int {
    pq := &PriorityQueue{}
    heap.Init(pq)
    count := [26]int{}
    for _, t := range tasks {
        count[t - 'A']++
    }

    for _, c := range count {
        if c == 0 {
            continue
        }
        t := Task{
            count: c,
            time: 0,
        }
        heap.Push(pq, t)
    }

    queue := []Task{}
    cycleTime := 0
    for pq.Len() > 0 || len(queue) > 0 {
        cycleTime++
        if pq.Len() > 0 {
            x := heap.Pop(pq)
            task := x.(Task)
            task.count--
            if task.count > 0 {
                task.time = cycleTime + n
                queue = append(queue, task)
            }
        }
        if len(queue) > 0 {
            front := queue[0]
            if front.time <= cycleTime {
                heap.Push(pq, front)
                queue = queue[1:]
            }
        }
    }

    return cycleTime
}