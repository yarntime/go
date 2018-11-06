func maximumSwap(num int) int {
    digit := make([]int, 0)
    t := num
    for t != 0 {
        digit = append([]int{t % 10}, digit...)
        t /= 10
    }
    sorted := make([]int, len(digit))
    copy(sorted, digit)
    
    sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
    
    pos := -1
    var a, b int
    for i := 0; i < len(sorted); i++ {
        if sorted[i] != digit[i] {
            pos = i
            a = sorted[i]
            b = digit[i]
            break
        }
    }
    
    if pos != -1 {
        digit[pos] = a
        for i := len(digit) - 1; i > pos; i-- {
            if digit[i] == a {
                digit[i] = b
                break
            }
        }
    }
    
    result := 0
    for i := 0; i < len(digit); i++ {
        result = result * 10 + digit[i]
    }
    return result
}