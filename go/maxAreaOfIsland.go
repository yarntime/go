var direct = [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}

func maxAreaOfIsland(grid [][]int) int {
    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 {
                continue
            }
            area := getArea(grid, i, j)
            if area > result {
                result = area
            }
        }
    }
    return result
}

func getArea(grid [][]int, i, j int) int {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
        return 0
    }
    if grid[i][j] == 0 {
        return 0
    }
    result := 1
    grid[i][j] = 0
    for k := 0; k < len(direct); k++ {
        result += getArea(grid, i + direct[k][0], j + direct[k][1])
    }
    return result
}