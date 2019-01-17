/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    inorder := helper(root)
    result := &TreeNode{}
    cur := result
    for i := 0; i < len(inorder); i++ {
        cur.Right = &TreeNode{
            Val: inorder[i],
            Left: nil,
            Right: nil,
        }
        cur = cur.Right
    }
    return result.Right
}

func helper(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    result := make([]int, 0)
    result = append(result, helper(root.Left)...)
    result = append(result, root.Val)
    result = append(result, helper(root.Right)...)
    return result
}