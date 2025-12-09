func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l, r := 0, len(people)-1
	var boats int

	for l <= r {
		boats++
		if people[l]+people[r] <= limit {
			l++
		}
		r--
	}

	return boats
}