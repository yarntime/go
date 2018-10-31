func trap(height []int) int {
    if len(height) <= 2 {
        return 0
    }
    result := 0
    current := 0
    stack := make([]int, 0)
    for current < len(height) {
        if len(stack) == 0 || height[stack[len(stack) - 1]] > height[current] {
            stack = append(stack, current)
            current++
        } else {
            t := stack[len(stack) - 1]
            stack = stack[0:len(stack) - 1]
            if len(stack) == 0 {
                continue
            }
            result += (min(height[current], height[stack[len(stack) - 1]]) - height[t]) * (current - stack[len(stack) - 1] - 1)
        }
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}