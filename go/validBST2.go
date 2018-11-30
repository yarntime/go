/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var pre *TreeNode

func isValidBST(root *TreeNode) bool {
    pre = nil
    return helper(root)
}

func helper(root *TreeNode) bool {
    if root == nil {
        return true
    }
    b := helper(root.Left)
    if !b {
        return b
    }
    if pre != nil {
        if root.Val <= (*pre).Val {
            return false
        }
    }
    pre = root
    return helper(root.Right)
}