package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	result = append(result, []int{root.Val})
	result = append(result, merge(levelOrder(root.Left), levelOrder(root.Right))...)
	return result
}

func merge(arr1, arr2 [][]int) [][]int {
	len1 := len(arr1)
	len2 := len(arr2)
	i := 0
	result := make([][]int, 0)
	for i < len1 && i < len2 {
		arr1[i] = append(arr1[i], arr2[i]...)
		result = append(result, arr1[i])
		i++
	}

	for i < len1 {
		result = append(result, arr1[i:]...)
		i++
	}

	for i < len2 {
		result = append(result, arr2[i:]...)
		i++
	}
	return result
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	r := levelOrder(tree)

	for i := 0; i < len(r); i++ {
		fmt.Printf("%v\n", r[i])
	}
}
