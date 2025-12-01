func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	l, r := 0, len(s)-1
	runeS := []rune(s)

	for l < r {
		charL := unicode.ToLower(runeS[l])
		if !unicode.IsDigit(charL) && !unicode.IsLetter(charL) {
			l++
			continue
		}

		charR := unicode.ToLower(runeS[r])
		if !unicode.IsDigit(charR) && !unicode.IsLetter(charR) {
			r--
			continue
		}

		if charL != charR {
			return false
		}
		l++
		r--
	}
	return true
}