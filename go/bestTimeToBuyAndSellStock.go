func maxProfit(prices []int) int {
    curMin := 1 << 31 - 1
    result := 0
    for i := 0; i < len(prices); i++ {
        result = max(prices[i] - curMin, result)
        curMin = min(curMin, prices[i])
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}