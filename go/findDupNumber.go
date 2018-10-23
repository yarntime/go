func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[abs(nums[i]) - 1] < 0 {
			return abs(nums[i])
		}
		nums[abs(nums[i]) - 1] = -nums[abs(nums[i]) - 1]
	}

	return -1
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}