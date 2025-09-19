func groupAnagrams(strs []string) [][]string {
    if len(strs) == 1 {
        return [][]string{strs}
    }
    found := make(map[string][]string)
    for _, s := range strs {
        runes := []rune(s)
        slices.Sort(runes)
        sortedKey := string(runes)
        found[sortedKey] = append(found[sortedKey], s)
    }

    var answer [][]string
    for _, val := range found {
        answer = append(answer, val)
    }
    return answer
}