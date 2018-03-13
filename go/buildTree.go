func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(preorder, inorder, 0, len(preorder), 0, len(inorder))
}

func helper(preorder []int, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart >= preEnd || inStart >= inEnd {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[preStart],
		Left:  nil,
		Right: nil,
	}
	pos := getLen(preorder[preStart], inorder, inStart, inEnd)
	root.Left = helper(preorder, inorder, preStart + 1, preStart + pos - inStart + 1, inStart, pos)
	root.Right = helper(preorder, inorder, preStart + pos - inStart + 1, preEnd, pos + 1, inEnd)
	return root
}

func getLen(val int, inorder []int, start, end int) int {
	for i := start; i < end; i++ {
		if inorder[i] == val {
			return i
		}
	}
	return 0
}