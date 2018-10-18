func singleNumber(nums []int) []int {
    firstOr := 0
    for i := 0; i < len(nums); i++ {
        firstOr ^= nums[i]
    }
    pos := getFirstBitOne(firstOr)
    
    first, second := 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] & (1 << pos) != 0 {
            first ^= nums[i]
        } else {
            second ^= nums[i]
        }
    }
    return []int{first, second}
}

func getFirstBitOne(a int) uint {
    pos := uint(0)
    for pos <= 32 {
        if a & (1 << pos) != 0 {
            break
        }
        pos++
    }
    return pos
}