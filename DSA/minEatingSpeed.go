func minEatingSpeed(piles []int, h int) int {
	maxElem := 0
	for _, val := range piles {
		if val > maxElem {
			maxElem = val
		}
	}

	l, r := 1, maxElem
	answer := 0

	for l <= r {
		mid := (l + r) / 2
		totalHours := 0
		for _, p := range piles {
			totalHours += (p + mid - 1) / mid
		}

		if totalHours <= h {
			answer = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return answer
}