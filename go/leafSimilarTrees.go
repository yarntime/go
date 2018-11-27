/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaf1 := new([]int)
    leaf2 := new([]int)
    leafValues(root1, leaf1)
    leafValues(root2, leaf2)
    
    if len(*leaf1) != len(*leaf2) {
        return false
    }

    for i := 0; i < len(*leaf1); i++ {
        if (*leaf1)[i] != (*leaf2)[i] {
            return false
        }
    }
    return true
}

func leafValues(root *TreeNode, leafs *[]int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *leafs = append(*leafs, root.Val)
        return
    }
    if root.Left != nil {
        leafValues(root.Left, leafs)
    }
    if root.Right != nil {
        leafValues(root.Right, leafs)
    }
}