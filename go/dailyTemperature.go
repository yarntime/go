func dailyTemperatures(T []int) []int {
    result := make([]int, len(T))
    stack := make([]int, 0)
    for i := 0; i < len(T); i++ {
        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            for len(stack) != 0 && T[i] > T[stack[len(stack) - 1]] {
                result[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
                stack = stack[:len(stack) - 1]
            }
            stack = append(stack, i)
        }
    }
    for len(stack) != 0 {
        result[stack[len(stack) - 1]] = 0
        stack = stack[:len(stack) - 1]
    }
    return result
}