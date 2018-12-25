func numFriendRequests(ages []int) int {
    counts := make([]int, 121)
    for _, age := range ages {
        counts[age]++
    }
    preSum := make([]int, 121)
    for i := 1; i < 120; i++ {
        preSum[i] += preSum[i - 1] + counts[i]
    }
    res := 0
    for i := 120; i >= 1; i-- {
        lower := i / 2 + 7
        if counts[i] == 0 || lower >= i {
            continue
        }
        res += counts[i] * (counts[i] - 1)
        res += counts[i] * (preSum[i - 1] - preSum[lower])
    }
    return res
}