func findLength(A []int, B []int) int {
    result := 0
    dp := make([][]int, 0)
    for i := 0; i <= len(A); i++ {
        dp = append(dp, make([]int, len(B) + 1))
    }
    for i := 1; i <= len(A); i++ {
        for j := 1; j <= len(B); j++ {
            if A[i - 1] == B[j - 1] {
                dp[i][j] = dp[i - 1][j - 1] + 1
            } else {
                dp[i][j] = 0
            }
            if dp[i][j] > result {
                result = dp[i][j]
            }
        }
    }
    return result
}
