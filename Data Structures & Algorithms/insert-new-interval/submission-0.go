func insert(intervals [][]int, newInterval []int) [][]int {
    merged := make([][]int, 0)
    newIntervalUsed := false
    i := 0
    for i < len(intervals) || !newIntervalUsed {
        var interval []int
        if !newIntervalUsed && (i >= len(intervals) || (i < len(intervals) && intervals[i][0] > newInterval[0])) {
            interval = newInterval
            newIntervalUsed = true
        } else {
            interval = intervals[i]
            i++
        }

        if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
            merged = append(merged, interval)
        } else {
            merged[len(merged)-1][1] = max(interval[1], merged[len(merged)-1][1])
        }
    }

    return merged
}
