func nextGreaterElements(nums []int) []int {
    newNums := make([]int, len(nums))
    copy(newNums, nums)
    newNums = append(newNums, nums...)
    mark := make([]int, len(newNums))
    stack := make([]int, 0)
    for i := 0; i < len(newNums); i++ {
        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            for len(stack) > 0 && newNums[i] > newNums[stack[len(stack) - 1]] {
                mark[stack[len(stack) - 1]] = newNums[i]
                stack = stack[0:len(stack) - 1]
            }
            stack = append(stack, i)
        }
    }
    for len(stack) != 0 {
        mark[stack[len(stack) - 1]] = -1
        stack = stack[0:len(stack) - 1]
    }
    result := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        result = append(result, mark[i])
    }
    return result
}