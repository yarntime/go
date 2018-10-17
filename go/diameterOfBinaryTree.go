/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    longest := 0
    helper(root, &longest)
    return longest - 1
}

func helper(root *TreeNode, longest *int) int {
    if root == nil {
        return 0
    }
    leftLongest := helper(root.Left, longest)
    rightLongest := helper(root.Right, longest)
    cur := leftLongest + rightLongest + 1
    if cur > *longest {
        *longest = cur
    }
    return max(leftLongest, rightLongest) + 1
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}