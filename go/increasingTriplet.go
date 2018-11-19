func increasingTriplet(nums []int) bool {
    dp := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
                if dp[i] >= 3 {
                    return true
                }
            }
        }
    }
    return false
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}