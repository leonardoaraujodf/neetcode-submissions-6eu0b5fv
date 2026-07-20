func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	half := (m + n + 1) / 2
	// i + j = (total + 1) / 2
	// j = (total+1)/2 - i
	left, right := 0, m
	for left <= right {
		i := left + (right-left)/2
		j := half - i
		Aright := math.MaxInt
		Aleft  := math.MinInt
		Bright := math.MaxInt
		Bleft := math.MinInt
		if i < m {
			Aright = nums1[i]
		}
		if i > 0 {
			Aleft = nums1[i-1]
		}
		if j < n {
			Bright = nums2[j]
		}
		if j > 0 {
			Bleft = nums2[j-1]
		}
		fmt.Println("Aright: ", Aright, "Aleft: ", Aleft, "Bright: ", Bright, "Bleft:", Bleft)
		if Aleft <= Bright && Bleft <= Aright {
			if (m+n) % 2 == 0 {
				return (float64(max(Aleft, Bleft)) + float64(min(Aright, Bright))) / 2.0
			}
			return float64(max(Aleft, Bleft))
		} else if Aleft > Bright {
			right = i - 1
		} else {
			left = i + 1
		}
	}

	return -1
}
