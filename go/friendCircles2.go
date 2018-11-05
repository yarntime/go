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
		if set[i] == i {
			result ++
		}
	}
	return result
}

func union(a, b int) {
	fa := find(a)
	fb := find(b)
	if fa != fb {
		set[fa] = fb
	}
}

func find(x int) int {
	if set[x] == x {
		return x
	}
	set[x] = find(set[x])
	return set[x]
}