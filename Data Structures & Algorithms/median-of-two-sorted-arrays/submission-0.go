func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A := nums1
	B := nums2

	if len(A) > len(B) {
		A, B = B, A
	}
	m := len(A)
	n := len(B)
	left := 0
	right := m
	
	for left <= right {
		i := left + (right - left) / 2
		j := (m + n + 1) / 2 - i
		Aleft := math.MinInt
		if i > 0 {
			Aleft = A[i - 1]
		}

		Aright := math.MaxInt
		if i < m {
			Aright = A[i]
		}

		Bleft := math.MinInt
		if j > 0 {
			Bleft = B[j - 1]
		}

		Bright := math.MaxInt
		if j < n {
			Bright = B[j]
		}

		if Aleft <= Bright && Bleft <= Aright {
			if (m+n) % 2 == 0 {
				return float64(max(Aleft, Bleft)+min(Aright, Bright)) / 2.0
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
