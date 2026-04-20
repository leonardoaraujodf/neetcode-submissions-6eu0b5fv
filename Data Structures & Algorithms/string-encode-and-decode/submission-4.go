type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return string("")
	}
	res := ""
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "#" + str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	i := 0
	res := []string{}
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		n, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:i+n])
		i += n
	}
	return res
}
