func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for !IsAlpha(rune(s[left])) {
			left++
			if left == right {
				return true
			}
		}

		c1 := unicode.ToLower(rune(s[left]))
		for !IsAlpha(rune(s[right])) && left < right {
			right--
			if left == right {
				return true
			}
		}

		c2 := unicode.ToLower(rune(s[right]))
		if c1 != c2 {
			return false
		}
		left++
		right--
	}
	return true
}

func IsAlpha(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}