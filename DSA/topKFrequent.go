func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int, len(nums))
	for _, val := range nums {
		count[val]++
	}

	pair := make([][]int, len(nums)+1)
	for num, freq := range count {
		pair[freq] = append(pair[freq], num)
	}

	var solution []int
	for i := len(pair) - 1; i >= 0 && len(solution) < k; i-- {
		for _, num := range pair[i] {
			if len(solution) >= k {
				break
			}
			solution = append(solution, num)
		}
	}
	return solution

}