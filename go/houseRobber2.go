func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    return max(getMax(nums, 0, len(nums) - 1), getMax(nums, 1, len(nums)))
}

func getMax(nums []int, start, end int) int {
    if len(nums) < 0 {
        return 0
    }
    pre1, pre2 := 0, 0
    for i := start; i < end; i++ {
        pre2, pre1 = pre1, max(pre2 + nums[i], pre1)
    }
    return pre1
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}