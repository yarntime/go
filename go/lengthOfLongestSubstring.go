func lengthOfLongestSubstring(str string) int {
    start := 0
    result := 0
    m := [256]int{}
    for i := range m {
        m[i] = -1
    }
    for i := 0; i < len(str); i++ {
        if m[str[i]] < start {
            if i - start + 1 > result {
                result = i - start + 1
            }
        } else {
            start = m[str[i]] + 1
        }
        m[str[i]] = i
    }
    return result
}