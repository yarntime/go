func canJump(nums []int) bool {
    next := 0
    for i := 0; i < len(nums); i++ {
        if next < i {
            return false
        }
        next = max(next, i + nums[i])
    }
    return true
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}