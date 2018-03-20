func addOneRow(root *TreeNode, v int, d int) *TreeNode {
    if d == 1 {
        return &TreeNode {
            Val: v,
            Left: root,
            Right: nil,
        }
    }
    return helper(root, 1, v, d)
}

func helper(root *TreeNode, curDepth, v, d int) *TreeNode {
    if root == nil {
        return nil
    }
    if curDepth + 1 == d {
        tmpLeftNode := &TreeNode {
            Val: v,
            Left: root.Left,
            Right: nil,
        }
        root.Left = tmpLeftNode
        tmpRightNode := &TreeNode {
            Val: v,
            Left: nil,
            Right: root.Right,
        }
        root.Right = tmpRightNode
        return root
    }
    root.Left = helper(root.Left, curDepth + 1, v, d)
    root.Right = helper(root.Right, curDepth + 1, v , d)
    return root
}