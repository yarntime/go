func sortColors(nums []int)  {
    colors := [3]int{}
    for i := 0; i < len(nums); i++ {
        colors[nums[i]]++
    }
    
    start := 0
    for i := 0; i < 3; i++ {
        for colors[i] > 0 {
            nums[start] = i
            colors[i]--
            start++
        }
    }
}