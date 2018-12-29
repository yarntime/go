func wiggleSort(nums []int)  {
    sort.Ints(nums)
    tmp := make([]int, len(nums))
    
    i, j := (len(nums) + 1) / 2 - 1, len(nums) - 1
    
    for cur := 0; cur < len(nums); cur ++ {
        if cur & 1 != 0 {
            tmp[cur] = nums[j]
            j--
        } else {
            tmp[cur] = nums[i]
            i--
        }
    }
    copy(nums, tmp)
}