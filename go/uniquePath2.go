func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    dp := [101][101]int{}
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    dp[1][1] = 1
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if obstacleGrid[i - 1][j - 1] == 1 {
                dp[i][j] = -1
            } else {
                if dp[i - 1][j] < 0 && dp[i][j - 1] < 0 {
                    dp[i][j] = -1
                }
                if dp[i - 1][j] > 0 {
                    dp[i][j] += dp[i - 1][j]
                }
                if dp[i][j - 1] > 0 {
                    dp[i][j] += dp[i][j - 1]
                }
            }
        }
    }
    if dp[m][n] < 0 {
        return 0
    }
    return dp[m][n]
}