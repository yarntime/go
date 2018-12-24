func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
func shortestToChar(S string, C byte) []int {
    length := len(S)
    ret := make([]int, length)
    index := []int{-length}
    for i := 0; i < length; i++ {
        if S[i] == C {
            index = append(index, i)
        }
    }
    index = append(index, length * 2)
    for i, j := 0, 0; i < length; i++ {
        if i < index[j+1] {
            ret[i] = Min(index[j+1]-i, i - index[j])
        }
        if i == index[j+1] {
            ret[i] = 0
            j++
        }
    }
    return ret
}
