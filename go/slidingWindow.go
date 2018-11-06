func maxSlidingWindow(nums []int, k int) []int {
    result := make([]int, 0)
    queue := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if len(queue) == 0 {
            queue = append(queue, i)
        } else {
            for len(queue) != 0 && nums[i] > nums[queue[len(queue) - 1]] {
                queue = queue[0:len(queue) - 1]
            }
            queue = append(queue, i)
        }
        if queue[len(queue) - 1] - queue[0] >= k {
            queue = queue[1:]
        }
        if i >= k - 1 {
            result = append(result, nums[queue[0]])
        }
    }
    return result
}