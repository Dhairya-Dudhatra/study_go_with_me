func mergeAlternately(word1 string, word2 string) string {

	answer := ""
	p1 := 0
	rune1 := []rune(word1)
	rune2 := []rune(word2)

	for p1 < len(rune1) && p1 < len(rune2) {
		answer += string(rune1[p1]) + string(rune2[p1])
		p1++
	}
	if p1 >= len(rune1) {
		answer += string(rune2[p1:])
	}
	if p1 >= len(rune2) {
		answer += string(rune1[p1:])
	}
	return answer
}