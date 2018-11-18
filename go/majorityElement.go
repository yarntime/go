func majorityElement(nums []int) int {
    cur := nums[0]
    sum := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == cur {
            sum++
        } else {
            sum--
            if sum == 0 {
                cur = nums[i]
                sum = 1
            }
        }
    }
    return cur
}