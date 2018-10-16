type StringSlice []string

func (self StringSlice) Len() int {
	return len(self)
}

func (self StringSlice) Less(i, j int) bool {
    a := self[i] + self[j]
    b := self[j] + self[i]
	return a > b
}

func (self StringSlice) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func largestNumber(nums []int) string {
    s  := StringSlice{}
    for i := 0; i < len(nums); i++ {
        s = append(s, strconv.Itoa(nums[i]))
    }
    sort.Sort(s)
    ans := ""
    for i := 0; i < len(s); i++ {
        ans += s[i]
    }
    for ans[0] == '0' && len(ans) > 1 {
        ans = ans[1:]
    }
    return ans
}