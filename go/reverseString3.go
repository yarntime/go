func reverseStr(s string, k int) string {
    b := []byte(s)
    end := len(b) - 1
    for i := 0; i < end; i += 2 * k {
        r := i + k - 1
        if end < r {
            r = end
        }
        for l := i; l < r; l, r = l+1, r-1 {
            b[l], b[r] = b[r], b[l]
        }
    }
    return string(b)
}