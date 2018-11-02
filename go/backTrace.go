var result [][]int

func combinationSum(candidates []int, target int) [][]int {
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
		result = append(result, tmp)
	}
	if curIndex >= len(candidates) {
		return
	}
	used[curIndex] = 1
	curSum += candidates[curIndex]
	helper(candidates, curIndex + 1, curSum, target, used)
	used[curIndex] = 0
	curSum -= candidates[curIndex]
	helper(candidates, curIndex + 1, curSum, target, used)
}