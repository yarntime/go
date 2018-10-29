func sortColors(nums []int)  {
    red, white, blue := 0, 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            red++
        } else if nums[i] == 1 {
            white++
        } else {
            blue++
        }
    }
    start := 0
    for red > 0 {
        nums[start] = 0
        start++
        red--
    }
    for white > 0 {
        nums[start] = 1
        start++
        white--
    }
    for blue > 0 {
        nums[start] = 2
        start++
        blue--
    }
}