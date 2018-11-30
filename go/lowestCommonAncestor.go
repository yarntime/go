type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode ,p *TreeNode, q *TreeNode) *TreeNode {
	lb := find(root.Left, p, q)
	rb := find(root.Right, p, q)
	if  lb && rb {
		return root
	}
	if lb {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if rb {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return nil
}

func find(root *TreeNode, p *TreeNode, q *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p || root == q {
		return true
	}
	return find(root.Left, p, q) || find(root.Right, p, q)
}
