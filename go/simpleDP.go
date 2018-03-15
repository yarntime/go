var dp [1000][1000]int
func longestPalindromeSubseq(s string) int {
    l := len(s)
    for i := 0; i < l; i++ {
        for j := 0; j < l; j++ {
            dp[i][i] = 0
        }
    }
    return helper(0, l - 1, s)
}

func helper(i, j int, s string) int {
    if i > j {
        return 0
    }
    if i == j {
        return 1
    }
    if dp[i][j] != 0 {
        return dp[i][j]
    }
    if s[i] == s[j] {
        dp[i][j] = helper(i + 1, j - 1, s) + 2
    } else {
        dp[i][j] = max(helper(i + 1, j, s), helper(i, j - 1, s))
    }
    return dp[i][j]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}