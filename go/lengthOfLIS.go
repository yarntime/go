func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    result := 1
    dp := make([]int, len(nums))
    dp[0] = 1
    for i := 1; i < len(nums); i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[j] + 1, dp[i])
            }
        }
        if dp[i] > result {
            result = dp[i]
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}