func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}
		if isPalindrome(s[l+1:r+1]) || isPalindrome(s[l:r]) {
			return true
		} else {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}