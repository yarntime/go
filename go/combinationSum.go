var result [][]int

func combinationSum(candidates []int, target int) [][]int {
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
			for j := 0; j < used[i]; j++ {
				tmp = append(tmp, candidates[i])
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
    flag := true
    for flag && curSum < target {
        used[curIndex] ++
	    curSum += candidates[curIndex]
	    helper(candidates, curIndex, curSum, target, used)
        flag = false
    }
    used[curIndex] --
	curSum -= candidates[curIndex]
	helper(candidates, curIndex + 1, curSum, target, used)
	
}