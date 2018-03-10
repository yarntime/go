func isValidBST(root *TreeNode) bool {
    tra := inorderTra(root)
    for i := 0; i < len(tra) - 1; i++ {
        if tra[i] >= tra[i + 1] {
            return false
        }
    }
    return true
}

func inorderTra(root *TreeNode) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    result = append(result, inorderTra(root.Left)...)
    result = append(result, root.Val)
    result = append(result, inorderTra(root.Right)...)
    return result
}