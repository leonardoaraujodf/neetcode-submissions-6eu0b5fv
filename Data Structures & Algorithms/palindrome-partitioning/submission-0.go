func partition(s string) [][]string {
	n := len(s)
	var result [][]string
	var curr []string
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			tmp := make([]string, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}

		for j := i; j < n; j++ {
			if !isPalindrome(s, i, j) {
				continue
			}
			curr = append(curr, s[i:j+1])
			backtrack(j + 1)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0)
	return result
}

func isPalindrome(s string, left, right int) bool {
	fmt.Println("s: ", s, "left: ", left, "right: ", right)
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
