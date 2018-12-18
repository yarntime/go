func validPalindrome(s string) bool {
    start := 0
    end := len(s) - 1
    for start < end {
        if s[start] != s[end] {
            return isPalindrome(s, start + 1, end) || isPalindrome(s, start, end - 1)
        }
        start++
        end--
    }
    return true
}

func isPalindrome(s string, start, end int) bool {
    for start < end {
        if s[start] != s[end] {
            return false
        }
        start++
        end--
    }
    return true
}