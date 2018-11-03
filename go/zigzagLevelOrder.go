/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    fromLeft := true
    for len(queue) != 0 {
        length := len(queue)
        level := make([]int, 0)
        for i := 0; i < length; i++ {
            front := queue[0]
            queue = queue[1:]
            if fromLeft {
                level = append(level, front.Val)
            } else {
                level = append([]int{front.Val}, level...)
            }
            if front.Left != nil {
                queue = append(queue, front.Left)
            }
            if front.Right != nil {
                queue = append(queue, front.Right)
            }
        }
        result = append(result, level)
        fromLeft = !fromLeft
    }
    return result
    
}