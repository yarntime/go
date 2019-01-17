func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        stri := strconv.Itoa(nums[i])
        strj := strconv.Itoa(nums[j])
        return stri + strj > strj + stri
    })
    result := ""
    for i := 0; i < len(nums); i++ {
        result += strconv.Itoa(nums[i])
    }
    cur := 0
    for cur < len(result) - 1 && result[cur] == '0' {
        cur++
    }
    return result[cur:]
}