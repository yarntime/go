func minPathSum(grid [][]int) int {
    m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if i == 1 {
                dp[i][j] += dp[i][j - 1] + grid[i - 1][j - 1]
            } else if j == 1 {
                dp[i][j] += dp[i - 1][j] + grid[i - 1][j - 1]
            } else {
                dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i - 1][j - 1]
            }
        }
    }
    return dp[m][n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}