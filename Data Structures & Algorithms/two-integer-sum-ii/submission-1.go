func twoSum(numbers []int, target int) []int {
   left := 0
   right := len(numbers)-1
   for left < right {
	  curr := numbers[left]+numbers[right]
	  if curr == target {
		return []int{left+1,right+1}
	  } else if curr > target {
		right--
	  } else if curr < target {
		left++
	  }
   }
   return []int{-1, -1}
}
