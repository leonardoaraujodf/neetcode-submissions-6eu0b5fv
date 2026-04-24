func characterReplacement(s string, k int) int {
	res := 0
	charSet := map[byte]bool{}
	for _, c := range s {
		charSet[byte(c)] = true
	}

	for c, _ := range charSet {
		l, count := 0, 0
		for r := 0; r < len(s); r++ {
			if s[r] == c {
				count++
			}

			for r - l + 1 - count > k {
				if s[l] == c {
					count--
				}
				l++
			}
			res = max(res, r - l + 1)
		}
	}
	
	return res
}
