func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    for len(nums) != 0 && nums[0] == nums[len(nums) - 1] {
        if nums[0] == target {
            return true
        }
        nums = nums[1:]
    }
    if len(nums) == 0 {
        return false
    }
    i := 0
    j := len(nums) - 1
    for i <= j {
        mid := i + (j - i) / 2
        if nums[mid] == target {
            return true
        } else if nums[mid] == nums[j] {
            j = mid - 1
        } else if nums[mid] <= nums[j] {
            if nums[mid] < target && target <= nums[j] {
                i = mid + 1
            } else {
                j = mid - 1
            }
        } else if nums[mid] >= nums[i] {
            if nums[i] <= target && target < nums[mid] {
                j = mid - 1
            } else {
                i = mid + 1
            }
        }
    }
    return false
}