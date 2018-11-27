/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Record struct {
    Node *TreeNode
    ID    int
}
func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    result := 0
    queue := make([]*Record, 0)
    queue = append(queue, &Record{Node: root, ID: 1})
    for len(queue) != 0 {
        size := len(queue)
        smallID := 0
        bigID := 0
        for i := 0; i < size; i++ {
            r := queue[0]
            queue = queue[1:]
            if i == 0 {
                smallID = r.ID
            }
            if i == size - 1 {
                bigID = r.ID
            }
            if r.Node.Left != nil {
                queue = append(queue, &Record{Node: r.Node.Left, ID: r.ID * 2})
            }
            if r.Node.Right != nil {
                queue = append(queue, &Record{Node: r.Node.Right, ID: r.ID * 2 + 1})
            }
        }
        if bigID - smallID > result {
            result = bigID - smallID
        }
    }
    return result + 1
}