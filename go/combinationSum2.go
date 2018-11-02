var result [][]int

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
	result = make([][]int, 0)
	used := make([]int, len(candidates))
	helper(candidates, 0, 0, target, used)
	return  result
}

func helper(candidates []int, curIndex int, curSum int, target int, used []int) {
	if curSum == target {
		tmp := make([]int, 0)
		for i := 0; i < len(used); i++ {
			if used[i] != 0 {
				tmp = append(tmp, candidates[i])
			}
		}
        for i := 0; i < len(result); i++ {
            if isSame(result[i], tmp) {
                return
            }
        }
		result = append(result, tmp)
        return
	}
	if curIndex >= len(candidates) {
		return
	}
    if curSum > target {
        return
    }
	used[curIndex] = 1
	curSum += candidates[curIndex]
	helper(candidates, curIndex + 1, curSum, target, used)
	used[curIndex] = 0
	curSum -= candidates[curIndex]
	helper(candidates, curIndex + 1, curSum, target, used)
}

func isSame(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}