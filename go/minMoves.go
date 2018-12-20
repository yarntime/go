func minMoves(nums []int) int {
    min := nums[0]
    for i := 1; i < len(nums); i++ {
        if min > nums[i] {
            min = nums[i]
        }
    }
    result := 0
    for i := 0; i < len(nums); i++ {
        result += nums[i] - min
    }
    return result
}