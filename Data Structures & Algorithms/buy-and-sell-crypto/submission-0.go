func maxProfit(prices []int) int {

	l := 0
	r := 1
	res := 0
	for r < len(prices) {
		curr := prices[r] - prices[l]
		res = max(res, curr)
		for curr < 0 && l < r {
			l++
			curr = prices[r] - prices[l]
		}
		r++
	}

	return res
}
