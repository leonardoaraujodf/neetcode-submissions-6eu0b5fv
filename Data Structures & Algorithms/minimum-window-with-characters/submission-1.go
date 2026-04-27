func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

    need := map[byte]int{}
	have := map[byte]int{}
	for _, c := range t {
		need[byte(c)]++
	}

	l := 0
	indexes := [3]int{math.MaxInt, math.MaxInt, math.MaxInt} // length, left, right
	for r := 0; r < len(s); r++ {
		c := s[r]
		have[c]++
		haveFreq := have[c]
		
		needFreq, ok := need[c]
		if !ok {
			continue
		}
		
		if needFreq != haveFreq {
			continue
		}

		for valid(need, have) && l <= r {
			length := r - l + 1
			fmt.Println(length)
			if length <= indexes[0] {
				indexes[0] = length
				indexes[1] = l
				indexes[2] = r
				fmt.Println(indexes)
			}

			c = s[l]
			have[c]--
			haveFreq = have[c]
			if haveFreq == 0 {
				delete(have, c)
			}
			l++
			fmt.Printf("c = %c, l = %d\n", c, l)
		}
	}

	if indexes[0] == math.MaxInt {
		return ""
	}
	res := s[indexes[1]:indexes[2]+1]
	return res
}

func valid(need map[byte]int, have map[byte]int) bool {
	for needKey, needVal := range need {
		haveVal, ok := have[needKey]
		if !ok {
			return false
		}
		if haveVal < needVal {
			return false
		}
	}
	fmt.Println("Here")
	return true
}
