func uniquePaths(m int, n int) int {
    dp := [101][101]int{}
    dp[1][1] = 1
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if i == 1 && j == 1 {
                continue
            }
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
        }
    }
    return dp[m][n]
}