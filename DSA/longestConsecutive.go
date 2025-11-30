func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	exist := make(map[int]bool)
	longestStreak := 0

	for _, num := range nums {
		exist[num] = true
	}

	for num := range exist {
		if !exist[num-1] {
			currentNum := num
			currentStreak := 1

			for exist[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}