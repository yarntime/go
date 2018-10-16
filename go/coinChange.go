var INF int = 1 << 31 - 1

func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    dp[0] = 0
    for i := 1; i <= amount; i++ {
        dp[i] = INF
        for j := 0; j < len(coins); j++ {
            if coins[j] <= i {
                dp[i] = min(dp[i - coins[j]] + 1, dp[i])
            }
        }
    }
    if dp[amount] == INF {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}