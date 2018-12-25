func largestRectangleArea(heights []int) int {
    heights = append(heights, 0)
    stack := make([]int, 0)
    result := 0
    for i := 0; i < len(heights); i++ {
        if len(stack) == 0 || heights[i] > heights[stack[len(stack) - 1]] {
            stack = append(stack, i)
            continue
        }
        for len(stack) != 0 && heights[i] <= heights[stack[len(stack) - 1]] {
            cur := stack[len(stack) - 1]
            stack = stack[0: len(stack) - 1]
            curSum := 0
            if len(stack) == 0 {
                curSum = heights[cur] * i
            } else {
                curSum = heights[cur] * (i - stack[len(stack) - 1] - 1)
            }
            if curSum > result {
                result = curSum
            }
        }
        stack = append(stack, i)
    }
    return result
}
