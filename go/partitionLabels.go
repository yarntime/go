func partitionLabels(s string) []int {
    lastPos := make([]int, 26)
    for i := len(s) - 1; i >= 0; i-- {
        if lastPos[s[i] - 'a'] == 0 {
            lastPos[s[i] - 'a'] = i
        }
    }
    result := make([]int, 0)
    strLen := 0
    maxPos := 0
    for i := 0; i < len(s); i++ {
        if i > maxPos {
            result = append(result, strLen)
            strLen = 0
            maxPos = 0
        } 
        maxPos = max(maxPos, lastPos[s[i] - 'a'])
        strLen++
    }
    result = append(result, strLen)
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}