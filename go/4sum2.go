func fourSumCount(A []int, B []int, C []int, D []int) int {
    length := len(A)
    abSum := make(map[int]int)
    for i := 0; i < length; i++ {
        for j := 0; j < length; j++ {
            abSum[A[i] + B[j]] += 1
        }
    }
    
    result := 0
    for i := 0; i < length; i++ {
        for j := 0; j < length; j++ {
            t := -C[i] - D[j]
            if c, ok := abSum[t]; ok {
                result += c
            }
        }
    }
    return result
}