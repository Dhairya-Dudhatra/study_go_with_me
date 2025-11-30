func subarraySum(nums []int, k int) int {
	prefixSumMap := make(map[int]int, len(nums)+1)
	var answer, prefixSum int
	prefixSumMap[0] = 1

	for _, num := range nums {
		prefixSum += num

		if val, ok := prefixSumMap[prefixSum-k]; ok {
			answer += val
		}
		prefixSumMap[prefixSum]++
	}
	return answer

}