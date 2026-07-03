func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevEnd := intervals[0][1]
	count := 0
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		// check overlapping
		if prevEnd > interval[0] {
			// one interval is going to overlaped.
			// be greedy and select the one which ends first.
			count++
			prevEnd = min(prevEnd, interval[1])
		} else {
			// update new interval ending time now
			prevEnd = interval[1]
		}
	}

	return count
}

// [1,2] [1,4] [2,4]