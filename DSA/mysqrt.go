func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 1, x/2
	ans := 0

	for l <= r {
		mid := l + (r-l)/2 // Idiomatic way to prevent (l+r) overflow

		// Use division to check: mid * mid <= x
		if mid <= x/mid {
			ans = mid   // Found a candidate, save it
			l = mid + 1 // Try to find a larger value
		} else {
			r = mid - 1 // mid is too big, look left
		}
	}
	return ans
}