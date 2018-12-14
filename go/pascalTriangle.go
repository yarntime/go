func generate(numRows int) [][]int {
    if numRows == 0 {
        return [][]int{}
    }
    result := make([][]int, 0)
    result = append(result, []int{1})
    for i := 1; i < numRows; i++ {
        result = append(result, make([]int, i + 1))
        result[i][0] = 1
        for j := 1; j < i; j++ {
            result[i][j] = result[i - 1][j - 1] + result[i - 1][j]
        }
        result[i][i] = 1
    }
    return result
}