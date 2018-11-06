/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    if root == nil {
        return 0
    }
    _, t := helper(root)
    return t
}

func helper(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    sl, tl := helper(root.Left)
    sr, tr := helper(root.Right)
    return root.Val + sl + sr, abs(sl - sr) + tl + tr
}

func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}