func findLongestChain(pairs [][]int) int {
    if len(pairs) == 0 {
        return 0
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][0] < pairs[j][0]
    })
    dp := make([]int, len(pairs))
    result := 1
    for i := 0; i < len(pairs); i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if pairs[j][1] < pairs[i][0] {
                dp[i] = max(dp[i], dp[j] + 1)
                result = max(result, dp[i])
            }
        }
    }
    return result
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}