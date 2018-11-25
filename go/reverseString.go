func reverseStr(s string, k int) string {
    var buffer bytes.Buffer
    n := len(s)
    
    for i := 0; i < n; i += 2 * k {
        maxIndex := Min(i + k, n)
        substr := Reverse(s[i : maxIndex])
        buffer.WriteString(substr)
        buffer.WriteString(s[maxIndex : Min(i + 2 * k, n)])
    }
    
    return buffer.String()
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}