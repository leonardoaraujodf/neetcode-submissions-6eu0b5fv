func majorityElement(nums []int) int {
    number, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			number = num
			count = 1
		} else {
			if num != number {
				count--
			} else {
				count++
			}
		}
	}
	return number
}
