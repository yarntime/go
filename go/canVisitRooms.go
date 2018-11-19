var marked []int

func canVisitAllRooms(rooms [][]int) bool {
    marked = make([]int, len(rooms))
    dfs(0, rooms)
    for i := 0; i < len(marked); i++ {
        if marked[i] == 0 {
            return false
        }
    }
    return true
}

func dfs(pos int, rooms [][]int) {
    if marked[pos] == 1 {
        return
    }
    marked[pos] = 1
    for i := 0; i < len(rooms[pos]); i++ {
        dfs(rooms[pos][i], rooms)
    }
}
