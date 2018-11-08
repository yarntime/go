/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var pre *TreeNode

func getMinimumDifference(root *TreeNode) int {
    if root == nil {
        return 0
    }
    result := 1 << 31 - 1
    pre = nil
    helper(root, &result)
    return result
}

func helper(root *TreeNode, result *int) {
    if root == nil {
        return
    }
    helper(root.Left, result)
    if pre != nil && root.Val - pre.Val < *result {
        *result = root.Val - pre.Val
    }
    pre = root
    helper(root.Right, result)
}