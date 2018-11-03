func subsets(nums []int) [][]int {
    total := 1
    length := 0
    for length < len(nums) {
        total = total << 1
        length++
    }
    result := make([][]int, 0)
    for i := 0; i < total; i++ {
        number := i
        cur := 0
        list := make([]int, 0)
        for number > 0 {
            if number & 1 != 0 {
                list = append(list, nums[cur])
            }
            cur++
            number = number >> 1
        }
        result = append(result, list)
    }
    return result
}