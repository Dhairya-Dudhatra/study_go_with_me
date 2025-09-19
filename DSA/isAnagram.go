

func isAnagram(s string, t string) bool {
    // If the lengths are different, they can't be anagrams.
    if len(s) != len(t) {
        return false
    }

    count := make(map[string]int, len(s))
    for i := range(len(s)) {
        count[s[i:i+1]]++ 
    }

    for i := range(len(t)) {
        if _, ok := count[t[i:i+1]]; !ok {
            return false
        } else {
            count[t[i:i+1]]--
        }

        if count[t[i:i+1]] == 0 {
            delete(count, t[i:i+1])
        }
    }
    return true
}