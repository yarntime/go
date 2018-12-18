func findDuplicates(nums []int) []int {
    result := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[abs(nums[i]) - 1] < 0 {
            result = append(result, abs(nums[i]))
        }
        nums[abs(nums[i]) - 1] = -nums[abs(nums[i]) - 1] 
    }
    return result
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}