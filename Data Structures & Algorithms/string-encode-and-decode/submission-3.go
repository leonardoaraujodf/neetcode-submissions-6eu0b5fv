type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// 1. Create a slice of runes to receive the resulted
	// encoded string
	// 2. Iterate over the slice of strings strs
	// 3. #{num} defines the maximum len of a string which is 200. We might have
	// a string starting with #xxx, where each x is a number.
	// 4. Define '/' as a escape character. If we have '#' inside the string, we encode
	// as '/#'.
	// 5. Also if we have a /, we encode as '//'
	if len(strs) == 0 {
		return string("")
	}
	var res []rune
	for _, str := range strs {
		res = append(res, '#')
		for _, c := range str {
			switch c {
			case '/', '#':
				res = append(res, '/')
			}
			res = append(res, c)
		}
	}
	fmt.Println(string(res))
	return string(res)
}

func (s *Solution) Decode(encoded string) []string {
	// 1. Iterate over the runes slice
	// 2. If we find a '#', it means it is a start of a string
	// 3. If we find a '\', skip to the next character and include
	// it as a character
	if len(encoded) == 0 {
		return []string{}
	}
	runes := []rune(encoded)
	var curr []rune
	var result []string
	init := true

	type state int
	const (
		processStr state = iota
		escapedChar
	)

	st := processStr

	for _, r := range runes {
		switch r {
		case '#':
			if st == processStr {
				if init {
					init = false
				} else {
					result = append(result, string(curr))
					curr = []rune{}
				}
				st = processStr
			} else if st == escapedChar {
				curr = append(curr, r)
				st = processStr
			}

			continue
		case '/':
			if st == escapedChar {
				curr = append(curr, r)
				st = processStr
			} else {
				st = escapedChar
				continue
			}
		default:
			if st == escapedChar {
				st = processStr
			}
			curr = append(curr, r)
		}
	}
	result = append(result, string(curr))
	return result
}
