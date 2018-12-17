/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Record struct {
    Node  *TreeNode
    Sn    int
}
func isCompleteTree(root *TreeNode) bool {
    if root == nil {
        return true
    }
    queue := make([]Record, 0)
    queue = append(queue, Record{Node: root, Sn: 1})
    previous := 0
    for len(queue) != 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            if node.Sn != previous + 1 {
                return false
            }
            previous = previous + 1
            if node.Node.Left != nil {
                queue = append(queue, Record{Node: node.Node.Left, Sn: node.Sn * 2})
            }
            if node.Node.Right != nil {
                queue = append(queue, Record{Node: node.Node.Right, Sn: node.Sn * 2 + 1})
            }
        }
    }
    return true
}