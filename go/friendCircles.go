var set []int

func findCircleNum(M [][]int) int {
    set = make([]int, len(M))
    for i := 0; i < len(M); i++ {
        set[i] = i
    }
    for i := 0; i < len(M) - 1; i++ {
        for j := i + 1; j < len(M); j++ {
            if M[i][j] == 1 {
                union(i, j)
            }
        }
    }
    result := 0
    for i := 0; i < len(M); i++ {
        if find(set[i]) == i {
            result ++
        }
    }
    return result
}

func union(a, b int) {
    fa := find(a)
    fb := find(b)
    if fa < fb {
        set[fb] = fa
    } else {
		set[fa] = fb
	}
}

func find(x int) int {
    for set[x] != x {
        x = set[x]
    }
    return x
}