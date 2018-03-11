func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePathHelper(root, "")
}

func binaryTreePathHelper(root *TreeNode, path string) []string {
	if root == nil {
        return []string{}
	}
	if root.Left == nil && root.Right == nil {
        return []string{path + strconv.Itoa(root.Val)}
	}
	result := []string{}
	newPath := path + strconv.Itoa(root.Val) + "->"
	result = append(result, binaryTreePathHelper(root.Left, newPath)...)
	result = append(result, binaryTreePathHelper(root.Right, newPath)...)
	return result
}