/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    i, e := traverse(root)
    return i + e
}

func traverse(t *TreeNode)(internal, external int){
    if t == nil {
        return 0,0
    }
    li, le := traverse(t.Left)
    ri, re := traverse(t.Right)
    return li+ri+abs(le)+abs(re), le + re + 1 - t.Val
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}