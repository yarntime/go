func twoSum(numbers []int, target int) []int {
    if len(numbers) < 2 {
        return []int{-1, -1}
    }
    for index, value := range numbers {
        pos := find(numbers, target - value)
        if pos != -1 && pos != index {
            return []int{index + 1, pos + 1}
        }
    }
    return []int{-1, -1}
}

func find(numbers []int, target int) int {
    i, j := 0, len(numbers)
    for i < j {
        mid := i + ((j - i) >> 1)
        if numbers[mid] == target {
            return mid
        } else if numbers[mid] > target {
            j = mid
        } else {
            i = mid + 1
        }
    }
    return -1
}