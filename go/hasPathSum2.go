func pathSum(root *TreeNode, sum int) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    newSum := sum - root.Val
    if newSum == 0 && root.Left == nil && root.Right == nil {
        result = append(result, []int{root.Val})
        return result
    }
    result = append(result, pathSum(root.Left, newSum)...)
    result = append(result, pathSum(root.Right, newSum)...)
    if len(result) != 0 {
        for i := 0; i < len(result); i++ {
            result[i] = append([]int{root.Val}, result[i]...)
        }
    }
    return result
}