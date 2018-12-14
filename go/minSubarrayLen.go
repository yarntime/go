func minSubArrayLen(s int, nums []int) int {
    var start, sum int
    min := math.MaxInt32
    for i, v := range nums {
        sum += v
        if sum >= s {
            for start < i && sum - nums[start] >= s {
                sum -= nums[start]
                start++
            }
            min = Min(i - start + 1, min)
        }
    }
    if min != math.MaxInt32 {
        return min
    }
    return 0
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}