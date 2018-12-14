func hasAlternatingBits(n int) bool {
    if n == 0 {
        return true
    }
    pre := n & 1
    n = n >> 1
    for n != 0 {
        c := n & 1
        if c == pre {
            return false
        }
        pre = c
        n = n >> 1
    }
    return true
}