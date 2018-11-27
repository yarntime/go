/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return helper(nums, 0, len(nums) - 1)
}

func helper(nums []int, start, end int) *TreeNode {
    if start > end {
        return nil
    }
    maxIndex := start
    maxVal := nums[start]
    for i := start + 1; i <= end; i++ {
        if nums[i] > maxVal {
            maxVal = nums[i]
            maxIndex = i
        }
    }
    root := &TreeNode{
        Val: maxVal,
    }
    root.Left = helper(nums, start, maxIndex - 1)
    root.Right = helper(nums, maxIndex + 1, end)
    return root
}