func integerBreak(n int) int {
    if n == 2 || n == 3 {
        return n - 1
    }
    result := 1
    for n > 4 {
        result *= 3
        n -= 3
    }
    return result * n
}