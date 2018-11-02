func nextGreaterElement(findNums []int, nums []int) []int {
    mark := make(map[int]int, len(nums))
    
    stack := make([]int, 0)
    
    for i := 0; i < len(nums); i++ {
        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            for len(stack) > 0 && nums[i] > nums[stack[len(stack) - 1]] {
                mark[nums[stack[len(stack) - 1]]] = nums[i]
                stack = stack[0:len(stack) - 1]
            }
            stack = append(stack, i)
        }
    }
    for len(stack) != 0 {
        mark[nums[stack[len(stack) - 1]]] = -1
        stack = stack[0:len(stack) - 1]
    }
    
    result := make([]int, 0)
    for i := 0; i < len(findNums); i++ {
        result = append(result, mark[findNums[i]])
    }
    return result
}