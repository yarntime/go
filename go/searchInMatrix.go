func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

    row := len(matrix)
	col := len(matrix[0])
	
    i := 0
    j := row * col - 1

	for i <= j {
        mid := i + (j - i) >> 1
        val := matrix[mid / col][mid % col]
        if val == target {
            return true
        } else if val > target {
            j = mid - 1
        } else {
            i = mid + 1
        }
	}

	return false
}