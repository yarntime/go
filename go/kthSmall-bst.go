/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    return Helper(root, &k)
}

func Helper(root *TreeNode, k *int) int {
    if root == nil {
        return -1
    }
    val := Helper(root.Left, k)
    if *k == 0 {
        return val
    }
    *k--
    if *k == 0 {
        return root.Val
    }
    return Helper(root.Right, k)
}