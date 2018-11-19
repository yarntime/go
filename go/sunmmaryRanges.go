func summaryRanges(nums []int) []string {
    result := make([]string, 0)
    if len(nums) <= 0 {
        return result
    }
    nums = append(nums, -1)
    curBegin := nums[0]
    curNum := nums[0]
    for i := 1; i < len(nums); i++ {
        if (nums[i] - curNum) != 1 {
            if curNum - curBegin == 0 {
                result = append(result, strconv.Itoa(curBegin))
            } else {
                result = append(result, strconv.Itoa(curBegin) + "->" + strconv.Itoa(curNum))
            }
            curBegin = nums[i]
        } 
        curNum = nums[i]
    }
    return result
}