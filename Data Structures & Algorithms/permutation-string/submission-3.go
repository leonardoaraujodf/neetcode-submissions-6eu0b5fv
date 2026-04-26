func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	m1 := [26]int{}
	m2 := [26]int{}
	for i := 0; i < len(s1); i++ {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if m1[i] == m2[i] {
			matches++
		}
	}

	left := 0
	for right := len(s1); right < len(s2); right++ {
		if matches == 26 {
			return true
		}

		idx := s2[right] - 'a'
		m2[idx]++
		if m1[idx] == m2[idx] {
			matches++
		} else if m1[idx] == m2[idx]-1 {
			matches--
		}

		idx = s2[left] - 'a'
		m2[idx]--
		if m1[idx] == m2[idx] {
			matches++
		} else if m1[idx] == m2[idx]+1 {
			matches--
		}
		left++
	}

	return matches == 26
}