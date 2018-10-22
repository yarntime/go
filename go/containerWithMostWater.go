func maxArea(height []int) int {
    i, j := 0, len(height) - 1
    result := 0
    for i < j {
        cur := min(height[i], height[j]) * (j - i)
        if cur > result {
            result = cur
        }
        if height[i] > height[j] {
            j--
        } else {
            i++
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