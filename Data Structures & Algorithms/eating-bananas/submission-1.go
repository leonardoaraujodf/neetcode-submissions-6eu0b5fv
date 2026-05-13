func minEatingSpeed(piles []int, h int) int {
	sum := 0
	for _, pile := range piles {
		sum += pile
	}
	l, r := 0, sum
	k := math.MaxInt
	for l <= r {
		mid := l + (r - l) / 2
		if check(mid, piles, h) {
			k = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return k
}

func check(k int, piles []int, h int) bool {
	if k == 0 {
		return false
	}

	// t, idx, curr := 0, 0, piles[0]
	t := 0
	for i := 0; i < len(piles); i++ {
		t += int(math.Ceil(float64(piles[i])/float64(k)))
	}
	//for idx < len(piles) {
	//  fmt.Println("idx: ", idx, "curr", curr, "t: ", t)
	//	curr = max(0, curr - k)
	//	if curr == 0 {
	//		idx++
	//		if idx < len(piles) {
	//			curr = piles[idx]
	//		}
	//	}
	//	t++
	//}
	// fmt.Println("k: ", k, "t: ", t, "h: ", h)
	return t <= h
}
