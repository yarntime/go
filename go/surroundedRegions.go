var dir = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var marked [][]int
func solve(board [][]byte)  {
    row := len(board)
    if row == 0 {
        return
    }
    col := len(board[0])
    if col == 0 {
        return
    }
    marked = make([][]int, row)
    for i := 0; i < row; i++ {
        marked[i] = make([]int, col)
    }
    
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if (i == 0 || j == 0 || i == row - 1 || j == col - 1) && board[i][j] == 'O' {
                helper(board, i, j, row, col)
            }
        }
    }
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if marked[i][j] != 1 {
                board[i][j] = 'X'
            }
        }
    }
}

func helper(board [][]byte, i, j, row, col int) {
    if i < 0 || i >= row || j < 0 || j >= col {
        return
    }
    if marked[i][j] != 0 {
        return
    }
    marked[i][j] = 2
    if board[i][j] != 'O' {
        return
    }
    marked[i][j] = 1
    for d := 0; d < 4; d++ {
        helper(board, i + dir[d][0], j + dir[d][1], row, col)
    }
}