func rob(nums []int) int {
    if len(nums) < 0 {
        return 0
    }
    pre1, pre2 := 0, 0
    for i := 0; i < len(nums); i++ {
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