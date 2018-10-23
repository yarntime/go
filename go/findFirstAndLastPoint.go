func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    first, second := -1, -1
    first = searchRangeHelper(nums, target, true)
    if first != -1 {
        second = searchRangeHelper(nums, target, false)
    }
    return []int{first, second}
}

func searchRangeHelper(nums []int, target int, first bool) int {
    	i, j := 0, len(nums)
	for i < j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			if first {
				j = mid
			} else {
				i = mid + 1
			}
		} else if nums[mid] > target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if first && i < len(nums) && nums[i] == target {
		return i
	} else if !first && i > 0 && nums[i - 1] == target {
		return i - 1
	}
	return -1
}