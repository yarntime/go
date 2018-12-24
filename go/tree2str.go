/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper(t *TreeNode, res *bytes.Buffer) {
    if t == nil {
        return
    }    
    res.WriteString(strconv.Itoa(t.Val))
    if t.Left == nil && t.Right == nil {
        return
    }
    if t.Right != nil {
        res.WriteString("(")
        helper(t.Left, res)
        res.WriteString(")")
        res.WriteString("(")
        helper(t.Right, res)
        res.WriteString(")")
    } else {
        res.WriteString("(")
        helper(t.Left, res)
        res.WriteString(")")
    }
    return
}

func tree2str(t *TreeNode) string {
    var res bytes.Buffer
    helper(t, &res)
    return res.String() 
}