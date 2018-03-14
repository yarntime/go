func sumOfLeftLeaves(root *TreeNode) int {
    return helper(root, 0)
}

func helper(root *TreeNode, prefix int) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return prefix * root.Val
    }
    return helper(root.Left, 1) + helper(root.Right, 0)
}