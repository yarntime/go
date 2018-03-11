package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumNumberHelper(root, 0)
}

func sumNumberHelper(root *TreeNode, v int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return v * 10 + root.Val
	}
	val := v * 10 + root.Val
	return sumNumberHelper(root.Left, val) + sumNumberHelper(root.Right, val)
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}
	fmt.Printf("%d\n", sumNumbers(tree))
}
