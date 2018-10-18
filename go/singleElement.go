func singleNonDuplicate(nums []int) int {
    start, end := 0, len(nums)
    for start < end - 1 {
        mid := (start + end) / 2
        if mid & 1 != 0 {
            mid -= 1
        }
        if nums[mid] != nums[mid + 1] {
            end = mid
        } else {
            start = mid + 2
        }
    }
    return nums[start]
}