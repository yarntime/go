func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    _, result := helper(root)
    return result
}

func helper(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    left, leftResult := helper(root.Left)
    right, rightResult := helper(root.Right)
    if leftResult && rightResult {
        left = left + 1
        right = right + 1
        diff := abs(left, right)
        if diff <= 1 {
            return max(left, right), true
        } else {
            return -1, false
        }
    } else {
        return -1, false
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a, b int) int {
    c := a - b
    if c < 0 {
        return -c
    }
    return c
}