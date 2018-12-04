type TreeNode struct {
	Val int
	Childs  []*TreeNode
}

func findLongestWay(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	for i := 0; i < len(root.Childs); i++ {
		cur := 0
		if root.Childs[i] != nil && root.Val == root.Childs[i].Val - 1 {
			cur = findLongestWay(root.Childs[i]) + 1
		} else {
			cur = findLongestWay(root.Childs[i])
		}
		if cur > result {
			result = cur
		}
	}
	return result
}