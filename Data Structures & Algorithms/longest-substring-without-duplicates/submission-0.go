func lengthOfLongestSubstring(s string) int {
	l, res := 0, 0
	m := map[byte]int{}
	for r := 0; r < len(s); r++ {
		m[s[r]]++
		for m[s[r]] > 1 {
			m[s[l]]--
			l++
		}
		res = max(res, r - l + 1)
	}
	return res
}
