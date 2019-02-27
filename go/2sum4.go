func findTarget(root *TreeNode, k int) bool {
    return traverse(root, root, k)
}

func traverse(root *TreeNode, curNode *TreeNode, k int) bool {
    remaining := k - curNode.Val
    if findNum(root, curNode.Val, remaining) {
        return true
    } else {
        if curNode.Left != nil {
            if traverse(root, curNode.Left, k) {
                return true
            }
        }
        if curNode.Right != nil {
            if traverse(root, curNode.Right, k) {
                return true
            }
        }
    }
    return false;
}

func findNum(root *TreeNode, except int, num int) bool {
    if root.Val == num && root.Val != except {
        return true
    } else if num < root.Val {
        if root.Left != nil {
            return findNum(root.Left, except, num)
        }
    } else {
        if root.Right != nil {
            return findNum(root.Right, except, num)
        }
    }
    return false
}