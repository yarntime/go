func nextGreaterElements(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    nums = append(nums, nums...)
    result := make([]int, len(nums))
    stack := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        result[i] = -1
        if len(stack) == 0 || nums[i] < nums[stack[len(stack) - 1]] {
            stack = append(stack, i)
        } else {
            for len(stack) != 0 && nums[stack[len(stack) - 1]] < nums[i] {
                result[stack[len(stack) - 1]] = nums[i]
                stack = stack[0: len(stack) - 1]
            }
            stack = append(stack, i)
        }
    }
    return result[0:len(nums)/2]
}