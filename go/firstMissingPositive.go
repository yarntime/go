func firstMissingPositive(nums []int) int {
    bigest := len(nums)
    for i := 0; i < bigest; i++ {
        for nums[i] > 0 && nums[i] <= bigest && nums[nums[i] - 1] != nums[i] {
            nums[nums[i] - 1], nums[i] = nums[i], nums[nums[i] - 1]
        } 
    }
    for i := 0; i < bigest; i++ {
        if i + 1 != nums[i] {
            return i + 1
        }
    }
    return bigest + 1
}