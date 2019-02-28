/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, result := helper(root)
    return result
}

func helper(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    leftHeight, isLeftBalance := helper(root.Left)
    if !isLeftBalance {
        return -1, false
    }
    rightHeight, isRightBalance := helper(root.Right)
    if !isRightBalance {
        return -1, false
    }
    return max(leftHeight, rightHeight) + 1, abs(leftHeight - rightHeight) <= 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}