func longestConsecutive(nums []int) int {
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
	}

	res := 0
	for _, num := range nums {
		if _, exists := m[num - 1]; exists {
			continue
		}

		// start of a sequence
		curr := num
		seq := 1
		for {
			curr += 1
			_, exists := m[curr]
			if !exists {
				break
			}
			seq += 1
		}
		res = max(res, seq)
	}

	return res
}
