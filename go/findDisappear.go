func findDisappearedNumbers(nums []int) []int {
    max := len(nums) + 1
    for _, num := range nums {
        nums[num % max - 1] += max
    }
    result := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] < max {
            result = append(result, i + 1)
        }
    }
    return result
}