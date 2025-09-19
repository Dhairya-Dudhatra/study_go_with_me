func longestCommonPrefix(strs []string) string {
    answer := strs[0]
    for _, str := range strs[1:] {
        var minLen int
        if len(answer) < len(str) {
            minLen = len(answer)
        } else {
            minLen = len(str)
        }
        temp := answer
        answer = ""
        for ind := range minLen {
            if temp[ind] != str[ind] {
                break
            }
            answer = temp[:ind+1]
        }
        if answer == "" {
            return answer
        }
    }
    return answer
}