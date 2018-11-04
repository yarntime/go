func getPos(start, end int, nums []int) int {
    i, j := start, end
    for i < j {
        for {
            j--
            if j <= start || nums[j] < nums[start] {
                break
            }
        }
        for {
            i++
            if i >= end || nums[i] > nums[start] {
                break
            }
        }
        if i < j {
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    nums[j], nums[start] = nums[start], nums[j]
    return j
}

func findKthLargest(nums []int, k int) int {
    start, end := 0, len(nums)
    length := len(nums)
    for {
        pos := getPos(start, end, nums)
        if length - pos == k {
            return nums[pos]
        } else if length - pos > k {
            start = pos + 1
        } else {
            end = pos
        }
    }
    return -1
}