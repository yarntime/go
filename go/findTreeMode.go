func findMode(root *TreeNode) []int {
	var res []int
	var max int
	traverse(root, &res, &max, make(map[int]int))
	return res
}

func traverse(root *TreeNode, res *[]int, max *int, m map[int]int) {
	if root == nil {
		return
	}
	traverse(root.Left, res, max, m)
	m[root.Val]++
	switch {
	case m[root.Val] == *max:
		*res = append(*res, root.Val)
	case m[root.Val] > *max:
		*max = m[root.Val]
		*res = []int{root.Val}
	}
	traverse(root.Right, res, max, m)
}