func productExceptSelf(nums []int) []int {

	output := make([]int, len(nums))
	output[0] = 1
	for i := 1; i < len(nums); i++ {
		output[i] = nums[i-1] * output[i-1]
	}
	prefixMul := 1
	for i := len(output) - 1; i >= 0; i-- {
		output[i] = prefixMul * output[i]
		prefixMul = prefixMul * nums[i]
	}
	return output
}