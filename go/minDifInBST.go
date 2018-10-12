package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var MAX_INT = 1<<31 - 1

func minDiffInBST(root *TreeNode) int {
	res, _ := helper(root, nil)
	return res
}

func helper(root, prev *TreeNode) (int, *TreeNode) {
	if root == nil {
		return MAX_INT, prev
	}
	lres, prev := helper(root.Left, prev)
	if prev != nil {
		lres = min(lres, abs(prev.Val-root.Val))
	}
	prev = root
	rres, prev := helper(root.Right, prev)
	res := min(lres, rres)
	return res, prev
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	tree := &TreeNode{
		Val: 90,
		Left: &TreeNode{
			Val: 69,
			Left: &TreeNode{
				Val:  49,
				Left: nil,
				Right: &TreeNode{
					Val:   52,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   89,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}
	fmt.Printf("%d\n", minDiffInBST(tree))
}
