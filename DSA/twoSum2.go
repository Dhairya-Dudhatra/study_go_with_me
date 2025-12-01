func twoSum(numbers []int, target int) []int {
	l, r := 1, len(numbers)

	for l < r {
		sum := numbers[l-1] + numbers[r-1]
		if sum == target {
			return []int{l, r}
		}
		if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}