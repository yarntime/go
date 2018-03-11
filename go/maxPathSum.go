package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var curMax int

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	curMax = math.MinInt32
	maxPathSumHelp(root)
	return curMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSumHelp(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLeft := max(0, maxPathSumHelp(root.Left))
	maxRight := max(0, maxPathSumHelp(root.Right))

	curMax = max(curMax, maxLeft + maxRight + root.Val)

	return max(maxLeft, maxRight) + root.Val
}

func main() {
	tree := &TreeNode{
		Val: -3,
		Left: &TreeNode{
			Val: -1,
			Left: nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Printf("%d\n", maxPathSum(tree))
}
