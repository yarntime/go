var direct = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    N := len(matrix)
    M := len(matrix[0])
    total := N * M
    cur := 0
    curPosX, curPosY := 0, 0
    result := make([]int, 0)
    for cur < total {
        if cur + 1 == total {
            result = append(result, matrix[curPosX][curPosY])
            break
        }
        l := [2]int{M - 1, N - 1}
        ld := 0
        for i := 0; i < 4; i++ {
            for j := 0; j < l[ld]; j++ {
                result = append(result, matrix[curPosX][curPosY])
                curPosX += direct[i][0]
                curPosY += direct[i][1]
                cur++
                if cur == total {
                    goto END
                }
            }
            ld = 1 - ld
        }
        curPosX++
        curPosY++
        M -= 2
        N -= 2
    }
    END: 
    return result
}