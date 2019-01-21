func leastBricks(wall [][]int) int {
    if len(wall) == 0 || len(wall[0]) == 0 {
        return 0
    }
    sums := make(map[int]int)
    max := 0
    for i := 0; i < len(wall); i++ {
        sum := 0
        for j := 0; j < len(wall[i]) - 1; j++ {
            sum += wall[i][j]
            sums[sum]++
            if sums[sum] > max {
                max = sums[sum]
            }
        }
    }
    return len(wall) - max
}