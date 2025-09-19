func twoSum(nums []int, target int) []int {
    found := make(map[int]int)
    for i, val := range nums {
        if index, ok := found[target-val]; ok {
            return []int{i, index}
        }
        found[val] = i
    }
    return nil
}