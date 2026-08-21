func letterCombinations(digits string) []string {
	keyboard := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{}
	curr := []byte{}
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(digits) {
			tmp := make([]byte, len(curr))
			copy(tmp, curr)
			res = append(res, string(tmp))
			return
		}

		for _, r := range keyboard[digits[idx]] {
			b := byte(r)
			curr = append(curr, b)
			backtrack(idx + 1)
			curr = curr[:len(curr)-1]
		}
	}
	if len(digits) > 0 {
		backtrack(0)
	}

	return res
}
