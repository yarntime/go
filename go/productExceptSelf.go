func productExceptSelf(nums []int) []int {
    mul := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        mul[i] = 1
    }
    curMulti := 1
    for i := 0; i < len(nums); i++ {
        mul[i] = curMulti
        curMulti *= nums[i]
    }
    curMulti = 1
    for i := len(nums) - 1; i >= 0; i-- {
        mul[i] *= curMulti
        curMulti = curMulti * nums[i]
    }
    return mul
}