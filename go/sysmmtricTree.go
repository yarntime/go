func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSysmmetric(root.Left, root.Right)
}

func isSysmmetric(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if (left == nil && right != nil) || (left != nil && right == nil) {
        return false
    }
    if left.Val == right.Val && isSysmmetric(left.Left, right.Right) && isSysmmetric(left.Right, right.Left) {
        return true
    }
    return false
}