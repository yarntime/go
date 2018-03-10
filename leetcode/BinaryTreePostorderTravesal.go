func postorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    result = append(result, postorderTraversal(root.Left)...)
    result = append(result, postorderTraversal(root.Right)...)
    result = append(result, root.Val)
    return result
}