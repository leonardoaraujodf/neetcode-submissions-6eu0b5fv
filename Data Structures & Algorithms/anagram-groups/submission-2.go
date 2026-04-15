func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, c := range str {
			count[c-'a']++
		}
		res[count] = append(res[count], str)
	}

	var ans [][]string
	for _, group := range res {
		ans = append(ans, group)
	}
	
	return ans
}
