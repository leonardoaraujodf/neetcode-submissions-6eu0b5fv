func groupAnagrams(strs []string) [][]string {
	var maps []map[rune]int
	var ans [][]string
	for _, str := range strs {
		// populate a map for the current string
		strMap := make(map[rune]int)
		for _, c := range str {
			strMap[c]++
		}
		
		if len(maps) == 0 {
			maps, ans = addNewAnagram(maps, ans, str, strMap)
		} else {
			isAnagram := false
			for mapIdx, m := range maps {
				if len(m) != len(strMap) {
					continue
				}
				found := true
				for strMapKey, strMapVal := range strMap {
					val, exists := m[strMapKey]
					if !exists || strMapVal != val {
						found = false
					}
				}
				if found {
					isAnagram = true
					ans[mapIdx] = append(ans[mapIdx], str)
					break
				}
			}
			if !isAnagram {
				// add a new anagram
				maps, ans = addNewAnagram(maps, ans, str, strMap)
			}
		}
	}

	return ans
}

func addNewAnagram(maps []map[rune]int, ans [][]string, str string, strMap map[rune]int) ([]map[rune]int, [][]string) {
	maps = append(maps, strMap)
	ans = append(ans, make([]string, 1))
	ans[len(ans) - 1][0] = str
	return maps, ans
}
