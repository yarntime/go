func reverseStr(s string, k int) string {
    result := make([]rune, 0)
    shouldReverse := true
    for i := 0; i < len(s); i += k {
        l := min(len(s) - i, k)
        if shouldReverse {
            result = append(result, reverse([]rune(s[i:i + l]))...)
        } else {
            result = append(result, []rune(s[i:i + l])...)
        }
        shouldReverse = !shouldReverse
    }
    return string(result)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func reverse(s []rune) []rune {
    for i := 0; i < len(s) / 2; i++ {
        s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
    }
    return s
}