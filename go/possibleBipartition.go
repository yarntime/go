func possibleBipartition(N int, dislikes [][]int) bool {
    set := make([]int, 0)
    for i := 0; i <= N; i++ {
        set = append(set, i)
    }
    for i := 0; i < len(dislikes); i++ {
        p1, p2 := dislikes[i][0], dislikes[i][1]
        if set[p2] == p2 {
            set[p2] = p1
        } else {
            l1, like1 := find(p1, set)
            l2, like2 := find(p2, set)
            if l1 == l2 && like1 == like2 {
                return false
            }
        }
    }
    return true
}

func find(p int, set []int) (int, int) {
    like := 0
    for set[p] != p {
        p = set[p]
        like = 1 - like
    }
    return p, like
}