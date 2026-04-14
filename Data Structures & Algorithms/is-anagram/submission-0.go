func isAnagram(s string, t string) bool {
    m1 := make(map[rune]int)
    m2 := make(map[rune]int)
    for _, val := range s {
        m1[val]++
    }
    for _, val := range t {
        m2[val]++
    }
    if len(m1) != len(m2) {
        return false
    }

    for key1, val1 := range m1 {
        val2, exists := m2[key1]
        if !exists || val1 != val2 {
            return false
        }
    }
    return true
}
