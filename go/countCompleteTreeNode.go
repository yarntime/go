/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return helper(root, 1)
}

func helper(root *TreeNode, sin int) int {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return sin
    }

    if root.Right == nil {
        return helper(root.Left, sin * 2)
    }
    return max(helper(root.Left, sin * 2), helper(root.Right, sin * 2 + 1))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}