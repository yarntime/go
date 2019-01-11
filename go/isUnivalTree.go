/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var value int
func isUnivalTree(root *TreeNode) bool {
    value = root.Val
    return helper(root)
}
func helper(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Val != value {
        return false
    }
    return helper(root.Left) && helper(root.Right)
}