/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    
    stack := make([]*TreeNode, 0)
    
    for root != nil || len(stack) != 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            root = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            result = append(result, root.Val)
            root = root.Right
        }
    }
    
    return result
}