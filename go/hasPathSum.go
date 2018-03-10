func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    newSum := sum - root.Val
    if newSum == 0 && root.Left == nil && root.Right == nil {
        return true
    }
    if hasPathSum(root.Left, newSum) || hasPathSum(root.Right, newSum) {
        return true
    }
    return false
}