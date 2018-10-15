func searchInsert(nums []int, target int) int {
    i := 0
    j := len(nums)
    for i < j {
        mid := (i + j) / 2
        if nums[mid] == target {
            return mid
        } else if (nums[mid] < target) {
            i = mid + 1
        } else {
            j = mid
        }
    }
    return i
}