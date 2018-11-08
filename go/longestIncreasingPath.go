var direct = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

var mark [][]int

func longestIncreasingPath(matrix [][]int) int {
	mark = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		mark[i] = make([]int, len(matrix[0]))
	}
	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			cur := dfs(matrix, -1, i, j)
			mark[i][j] = cur
			result = max(cur, result)
		}
	}
	return result
}

func dfs(matrix [][]int, pre, i, j int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return 0
	}
	if matrix[i][j] <= pre {
		return 0
	}
	if mark[i][j] != 0 {
		return mark[i][j]
	}
	result := 0
	for dir := 0; dir < 4; dir++ {
		result = max(result, dfs(matrix, matrix[i][j], i + direct[dir][0], j + direct[dir][1]))
	}
	mark[i][j] = result + 1
	return result + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}