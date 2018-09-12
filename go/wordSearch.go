var direct = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
    used := make([][]bool, len(board))
    for k := 0; k < len(board); k++ {
        used[k] = make([]bool, len(board[k]))
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if helper(board, used, word, 0, i, j) {
                return true
            }
        }
    }
    return false
}

func helper(board [][]byte, used [][]bool, word string, cur int, i int, j int) bool {
    if cur >= len(word) {
        return true
    }
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || used[i][j] == true {
        return false
    }
    if word[cur] != board[i][j] {
        return false
    }
    used[i][j] = true
	for d := 0; d < 4; d++ {
		if helper(board, used, word, cur + 1, i + direct[d][0], j + direct[d][1]) {
			used[i][j] = false
			return true
		}
	}
    used[i][j] = false
	return false
}